---
name: Javascript API
category: 
---

# Web3 JavaScript API

To connect to Daxxcoin node, you can use the `web3` object provided by the [web3.js library]. Under the hood it communicates to a local node through [RPC calls].

## Using callbacks

As this API is designed to work with a local RPC node and all its functions are by default use synchronous HTTP requests.con

If you want to make asynchronous request, you can pass an optional callback as the last parameter to most functions.
All callbacks are using an [error first callback](http://fredkschott.com/post/2014/03/understanding-error-first-callbacks-in-node-js/) style:

```js
web3.eth.getBlock(48, function(error, result){
    if(!error)
        console.log(result)
    else
        console.error(error);
})
```

## Batch requests

Batch requests allow queuing up requests and processing them at once.

```js
var batch = web3.createBatch();
batch.add(web3.eth.getBalance.request('0x0000000000000000000000000000000000000000', 'latest', callback));
batch.add(web3.eth.contract(abi).at(address).balance.request(address, callback2));
batch.execute();
```

## A note on big numbers in web3.js

You will always get a BigNumber object for balance values as JavaScript is not able to handle big numbers correctly.
Look at the following examples:

```js
"101010100324325345346456456456456456456"
// "101010100324325345346456456456456456456"
101010100324325345346456456456456456456
// 1.0101010032432535e+38
```

web3.js depends on the [BigNumber Library](https://github.com/MikeMcl/bignumber.js/) and adds it automatically.

```js
var balance = new BigNumber('131242344353464564564574574567456');
// or var balance = web3.eth.getBalance(someAddress);

balance.plus(21).toString(10); // toString(10) converts it to a number string
// "131242344353464564564574574567477"
```

The next example wouldn't work as we have more than 20 floating points, therefore it is recommended to keep you balance always in *wei* and only transform it to other units when presenting to the user:
```js
var balance = new BigNumber('13124.234435346456466666457455567456');

balance.plus(21).toString(10); // toString(10) converts it to a number string, but can only show max 20 floating points 
// "13145.23443534645646666646" // you number would be cut after the 20 floating point
```

## Web3 Javascript API Reference

### Usage

#### web3
The `web3` object provides all methods.

##### Example

```js
var Web3 = require('web3');
// create an instance of web3 using the HTTP provider.
var web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
```

***

#### web3.version.api

```js
web3.version.api
// or async
web3.version.getApi(callback(error, result){ ... })
```

##### Returns

`String` - The js api version.

##### Example

```js
var version = web3.version.api;
console.log(version); // "0.2.0"
```

***

#### web3.version.node

    web3.version.node
    // or async
    web3.version.getClient(callback(error, result){ ... })


##### Returns

`String` - The client/node version.

##### Example

```js
var version = web3.version.node;
console.log(version); 
```

***

#### web3.version.network

    web3.version.network
    // or async
    web3.version.getNetwork(callback(error, result){ ... })


##### Returns

`String` - The network protocol version.

##### Example

```js
var version = web3.version.network;
console.log(version); // 54
```

***

#### web3.isConnected

    web3.isConnected()

Should be called to check if a connection to a node exists

##### Parameters
none

##### Returns

`Boolean`

##### Example

```js
if(!web3.isConnected()) {
  
   // show some dialog to ask the user to start a node

} else {
 
   // start web3 filters, calls, etc
  
}
```

***

#### web3.setProvider

    web3.setProvider(provider)

Should be called to set provider.

##### Parameters
none

##### Returns

`undefined`

##### Example

```js
web3.setProvider(new web3.providers.HttpProvider('http://localhost:8545'));
```

***

#### web3.toHex

    web3.toHex(mixed);
 
Converts any value into HEX.

##### Parameters

1. `String|Number|Object|Array|BigNumber` - The value to parse to HEX. If its an object or array it will be `JSON.stringify` first. If its a BigNumber it will make it the HEX value of a number.

##### Returns

`String` - The hex string of `mixed`.

##### Example

```js
var str = web3.toHex({test: 'test'});
console.log(str); // '0x7b2274657374223a2274657374227d'
```

***

#### web3.toAscii

    web3.toAscii(hexString);

Converts a HEX string into a ASCII string.

##### Parameters

1. `String` - A HEX string to be converted to ascii.

##### Returns

`String` - An ASCII string made from the given `hexString`.

##### Example

```js
var str = web3.toAscii("0x657468657265756d000000000000000000000000000000000000000000000000");
console.log(str); // 
```

***

#### web3.fromAscii

    web3.fromAscii(string [, padding]);

Converts any ASCII string to a HEX string.

##### Parameters

1. `String` - An ASCII string to be converted to HEX.
2. `Number` - The number of bytes the returned HEX string should have. 

##### Returns

`String` - The converted HEX string.

##### Example

```js
var str = web3.fromAscii("js")
"0x6a73"

```

***

#### web3.toDecimal

    web3.toDecimal(hexString);

Converts a HEX string to its number representation.

##### Parameters

1. `String` - An HEX string to be converted to a number.


##### Returns

`Number` - The number representing the data `hexString`.

##### Example

```js
var number = web3.toDecimal('0x15');
console.log(number); // 21
```

***

#### web3.fromDecimal

    web3.fromDecimal(number);

Converts a number or number string to its HEX representation.

##### Parameters

1. `Number|String` - A number to be converted to a HEX string.

##### Returns

`String` - The HEX string representing of the given `number`.

##### Example

```js
var value = web3.fromDecimal('21');
console.log(value); // "0x15"
```

***

#### web3.fromWei

    web3.fromWei(number, unit)

Converts a number of wei into the following units:

- `kwei`/`ada`
- `mwei`/`babbage`
- `gwei`/`shannon`
- `szabo`
- `finney`
- `ether`
- `kether`/`grand`/`einstein`
- `mether`
- `gether`
- `tether`

##### Parameters

1. `Number|String|BigNumber` - A number or BigNumber instance.
2. `String` - One of the above units.


##### Returns

`String|BigNumber` - Either a number string, or a BigNumber instance, depending on the given `number` parameter.

##### Example

```js
var value = web3.fromWei('21000000000000', 'finney');
console.log(value); // "0.021"
```

***

#### web3.toWei

    web3.toWei(number, unit)

Converts an unit into wei. Possible units are:

- `kwei`/`ada`
- `mwei`/`babbage`
- `gwei`/`shannon`
- `szabo`
- `finney`
- `ether`
- `kether`/`grand`/`einstein`
- `mether`
- `gether`
- `tether`

##### Parameters

1. `Number|String|BigNumber` - A number or BigNumber instance.
2. `String` - One of the above units.

##### Returns

`String|BigNumber` - Either a number string, or a BigNumber instance, depending on the given `number` parameter.

##### Example

```js
var value = web3.toWei('1', 'ether');
console.log(value); // "1000000000000000000"
```

***

#### web3.toBigNumber

    web3.toBigNumber(numberOrHexString);

Converts a given number into a BigNumber instance.

See the [note on BigNumber](#a-note-on-big-numbers-in-javascript).

##### Parameters

1. `Number|String` - A number, number string or HEX string of a number.


##### Returns

`BigNumber` - A BigNumber instance representing the given value.


##### Example

```js
var value = web3.toBigNumber('200000000000000000000001');
console.log(value); // instanceOf BigNumber
console.log(value.toNumber()); // 2.0000000000000002e+23
console.log(value.toString(10)); // '200000000000000000000001'
```

***

### web3.net

#### web3.net.listening

    web3.net.listening
    // or async
    web3.net.getListening(callback(error, result){ ... })

This property is read only and says whether the node is actively listening for network connections or not.

##### Returns

`Boolean` - `true` if the client is actively listening for network connections, otherwise `false`.

##### Example

```js
var listening = web3.net.listening;
console.log(listening); // true of false
```

***

#### web3.net.peerCount

    web3.net.peerCount
    // or async
    web3.net.getPeerCount(callback(error, result){ ... })

This property is read only and returns the number of connected peers.

##### Returns

`Number` - The number of peers currently connected to the client.

##### Example

```js
var peerCount = web3.net.peerCount;
console.log(peerCount); // 4
```

***

### web3.eth

Contains the blockchain related methods.

##### Example

```js
var eth = web3.eth;
```

***

#### web3.eth.defaultAccount

    web3.eth.defaultAccount

This default address is used for the following methods (optionally you can overwrite it by specifying the `from` property):

- [web3.eth.sendTransaction()](#web3ethsendtransaction)
- [web3.eth.call()](#web3ethcall)

##### Values

`String`, 20 Bytes - Any address you own, or where you have the private key for.

*Default is* `undefined`.

##### Returns

`String`, 20 Bytes - The currently set default address.

##### Example

```js
var defaultAccount = web3.eth.defaultAccount;
console.log(defaultAccount); // ''

// set the default block
web3.eth.defaultAccount = '0x8888f1f195afa192cfee860698584c030f4c9db1';
```

***

#### web3.eth.defaultBlock

    web3.eth.defaultBlock

This default block is used for the following methods (optionally you can overwrite the defaultBlock by passing it as the last parameter):

- [web3.eth.getBalance()](#web3ethgetbalance)
- [web3.eth.getCode()](#web3ethgetcode)
- [web3.eth.getTransactionCount()](#web3ethgettransactioncount)
- [web3.eth.getStorageAt()](#web3ethgetstorageat)
- [web3.eth.call()](#web3ethcall)

##### Values

Default block parameters can be one of the following:

- `Number` - a block number
- `String` - `"earliest"`, the genisis block
- `String` - `"latest"`, the latest block (current head of the blockchain)
- `String` - `"pending"`, the currently mined block (including pending transactions)

*Default is* `latest`

##### Returns

`Number|String` - The default block number to use when querying a state.

##### Example

```js
var defaultBlock = web3.eth.defaultBlock;
console.log(defaultBlock); // 'latest'

// set the default block
web3.eth.defaultBlock = 231;
```

***

#### web3.eth.syncing

    web3.eth.syncing
    // or async
    web3.eth.getSyncing(callback(error, result){ ... })

This property is read only and returns the either a sync object, when the node is syncing or `false`.

##### Returns

`Object|Boolean` - A sync object as follows, when the node is currently syncing or `false`:
   - `startingBlock`: `Number` - The block number where the sync started.
   - `currentBlock`: `Number` - The block number where at which block the node currently synced to already.
   - `highestBlock`: `Number` - The estimated block number to sync to.

##### Example

```js
var sync = web3.eth.syncing;
console.log(sync);
/*
{
   startingBlock: 300,
   currentBlock: 312,
   highestBlock: 512
}
*/
```

***

#### web3.eth.isSyncing

    web3.eth.isSyncing(callback);

This convenience function calls the `callback` everytime a sync starts, updates and stops.

##### Returns

`Object` - a isSyncing object with the following methods:

  * `syncing.addCallback()`: Adds another callback, which will be called when the node starts or stops syncing.
  * `syncing.stopWatching()`: Stops the syncing callbacks.

##### Callback return value

- `Boolean` - The callback will be fired with `true` when the syncing starts and with `false` when it stopped.
- `Object` - While syncing it will return the syncing object:
   - `startingBlock`: `Number` - The block number where the sync started.
   - `currentBlock`: `Number` - The block number where at which block the node currently synced to already.
   - `highestBlock`: `Number` - The estimated block number to sync to.


##### Example

```js
web3.eth.isSyncing(function(error, sync){
    if(!error) {
        // stop all app activity
        if(sync === true) {
           // we use `true`, so it stops all filters, but not the web3.eth.syncing polling
           web3.reset(true);
        
        // show sync info
        } else if(sync) {
           console.log(sync.currentBlock);
        
        // re-gain app operation
        } else {
            // run your app init function...
        }
    }
});
```

***

#### web3.eth.coinbase

    web3.eth.coinbase
    // or async
    web3.eth.getCoinbase(callback(error, result){ ... })

This property is read only and returns the coinbase address were the mining rewards go to.

##### Returns

`String` - The coinbase address of the client.

##### Example

```js
var coinbase = web3.eth.coinbase;
console.log(coinbase); // "0x407d73d8a49eeb85d32cf465507dd71d507100c1"
```

***

#### web3.eth.mining

    web3.eth.mining
    // or async
    web3.eth.getMining(callback(error, result){ ... })


This property is read only and says whether the node is mining or not.


##### Returns

`Boolean` - `true` if the client is mining, otherwise `false`.

##### Example

```js
var mining = web3.eth.mining;
console.log(mining); // true or false
```

***

#### web3.eth.hashrate

    web3.eth.hashrate
    // or async
    web3.eth.getHashrate(callback(error, result){ ... })

This property is read only and returns the number of hashes per second that the node is mining with.


##### Returns

`Number` - number of hashes per second.

##### Example

```js
var hashrate = web3.eth.hashrate;
console.log(hashrate); // 493736
```

***

#### web3.eth.gasPrice

    web3.eth.gasPrice
    // or async
    web3.eth.getGasPrice(callback(error, result){ ... })


This property is read only and returns the current gas price.
The gas price is determined by the x latest blocks median gas price.

##### Returns

`BigNumber` - A BigNumber instance of the current gas price in wei.

See the [note on BigNumber](#a-note-on-big-numbers-in-javascript).

##### Example

```js
var gasPrice = web3.eth.gasPrice;
console.log(gasPrice.toString(10)); // "10000000000000"
```

***

#### web3.eth.accounts

    web3.eth.accounts
    // or async
    web3.eth.getAccounts(callback(error, result){ ... })

This property is read only and returns a list of accounts the node controls.

##### Returns

`Array` - An array of addresses controlled by client.

##### Example

```js
var accounts = web3.eth.accounts;
console.log(accounts); // ["0x407d73d8a49eeb85d32cf465507dd71d507100c1"] 
```

***

#### web3.eth.blockNumber

    web3.eth.blockNumber
    // or async
    web3.eth.getBlockNumber(callback(error, result){ ... })

This property is read only and returns the current block number.

##### Returns

`Number` - The number of the most recent block.

##### Example

```js
var number = web3.eth.blockNumber;
console.log(number); // 2744
```

***

#### web3.eth.getBalance

    web3.eth.getBalance(addressHexString [, defaultBlock] [, callback])

Get the balance of an address at a given block.

##### Parameters

1. `String` - The address to get the balance of.
2. `Number|String` - (optional) If you pass this parameter it will not use the default block set with [web3.eth.defaultBlock](#web3ethdefaultblock).
3. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

##### Returns

`String` - A BigNumber instance of the current balance for the given address in wei.

See the [note on BigNumber](#a-note-on-big-numbers-in-javascript).

##### Example

```js
var balance = web3.eth.getBalance("0x407d73d8a49eeb85d32cf465507dd71d507100c1");
console.log(balance); // instanceof BigNumber
console.log(balance.toString(10)); // '1000000000000'
console.log(balance.toNumber()); // 1000000000000
```

***

#### web3.eth.getBlock

     web3.eth.getBlock(blockHashOrBlockNumber [, returnTransactionObjects] [, callback])

Returns a block matching the block number or block hash.

##### Parameters

1. `String|Number` - The block number or hash. Or the string `"earliest"`, `"latest"` or `"pending"` as in the [default block parameter](#web3ethdefaultblock).
2. `Boolean` - (optional, default `false`) If `true`, the returned block will contain all transactions as objects, if `false` it will only contains the transaction hashes.
3. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

##### Returns

`Object` - The block object:

  - `number`: `Number` - the block number. `null` when its pending block.
  - `hash`: `String`, 32 Bytes - hash of the block. `null` when its pending block.
  - `parentHash`: `String`, 32 Bytes - hash of the parent block.
  - `nonce`: `String`, 8 Bytes - hash of the generated proof-of-work. `null` when its pending block.
  - `sha3Uncles`: `String`, 32 Bytes - SHA3 of the uncles data in the block.
  - `logsBloom`: `String`, 256 Bytes - the bloom filter for the logs of the block. `null` when its pending block.
  - `transactionsRoot`: `String`, 32 Bytes - the root of the transaction trie of the block
  - `stateRoot`: `String`, 32 Bytes - the root of the final state trie of the block.
  - `miner`: `String`, 20 Bytes - the address of the beneficiary to whom the mining rewards were given.
  - `difficulty`: `BigNumber` - integer of the difficulty for this block.
  - `totalDifficulty`: `BigNumber` - integer of the total difficulty of the chain until this block.
  - `extraData`: `String` - the "extra data" field of this block.
  - `size`: `Number` - integer the size of this block in bytes.
  - `gasLimit`: `Number` - the maximum gas allowed in this block.
  - `gasUsed`: `Number` - the total used gas by all transactions in this block.
  - `timestamp`: `Number` - the unix timestamp for when the block was collated.
  - `transactions`: `Array` - Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter.
  - `uncles`: `Array` - Array of uncle hashes.

##### Example

```js
var info = web3.eth.block(3150);
console.log(info);
/*
{
  "number": 3,
  "hash": "0xef95f2f1ed3ca60b048b4bf67cde2195961e0bba6f70bcbea9a2c4e133e34b46",
  "parentHash": "0x2302e1c0b972d00932deb5dab9eb2982f570597d9d42504c05d9c2147eaf9c88",
  "nonce": "0xfb6e1a62d119228b",
  "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
  "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  "transactionsRoot": "0x3a1b03875115b79539e5bd33fb00d8f7b7cd61929d5a3c574f507b8acf415bee",
  "stateRoot": "0xf1133199d44695dfa8fd1bcfe424d82854b5cebef75bddd7e40ea94cda515bcb",
  "miner": "0x8888f1f195afa192cfee860698584c030f4c9db1",
  "difficulty": BigNumber,
  "totalDifficulty": BigNumber,
  "size": 616,
  "extraData": "0x",
  "gasLimit": 3141592,
  "gasUsed": 21662,
  "timestamp": 1429287689,
  "transactions": [
    "0x9fc76417374aa880d4449a1f7f31ec597f00b1f6f3dd2d66f4c9c6c445836d8b"
  ],
  "uncles": []
}
*/
```

***

#### web3.eth.getBlockTransactionCount

    web3.eth.getBlockTransactionCount(hashStringOrBlockNumber [, callback])

Returns the number of transaction in a given block.

##### Parameters

1. `String|Number` - The block number or hash. Or the string `"earliest"`, `"latest"` or `"pending"` as in the [default block parameter](#web3ethdefaultblock).
2. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

##### Returns

`Number` - The number of transactions in the given block.

##### Example

```js
var number = web3.eth.getBlockTransactionCount("0x407d73d8a49eeb85d32cf465507dd71d507100c1");
console.log(number); // 1
```

***

##### web3.eth.getTransaction

    web3.eth.getTransaction(transactionHash [, callback])

Returns a transaction matching the given transaction hash.

##### Parameters

1. `String` - The transaction hash.
2. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.


##### Returns

`Object` - A transaction object its hash `transactionHash`:

  - `hash`: `String`, 32 Bytes - hash of the transaction.
  - `nonce`: `Number` - the number of transactions made by the sender prior to this one.
  - `blockHash`: `String`, 32 Bytes - hash of the block where this transaction was in. `null` when its pending.
  - `blockNumber`: `Number` - block number where this transaction was in. `null` when its pending.
  - `transactionIndex`: `Number` - integer of the transactions index position in the block. `null` when its pending.
  - `from`: `String`, 20 Bytes - address of the sender.
  - `to`: `String`, 20 Bytes - address of the receiver. `null` when its a contract creation transaction.
  - `value`: `BigNumber` - value transferred in Wei.
  - `gasPrice`: `BigNumber` - gas price provided by the sender in Wei.
  - `gas`: `Number` - gas provided by the sender.
  - `input`: `String` - the data sent along with the transaction.


##### Example

```js
var blockNumber = 668;
var indexOfTransaction = 0

var transaction = web3.eth.getTransaction(blockNumber, indexOfTransaction);
console.log(transaction);
/*
{
  "hash": "0x9fc76417374aa880d4449a1f7f31ec597f00b1f6f3dd2d66f4c9c6c445836d8b",
  "nonce": 2,
  "blockHash": "0xef95f2f1ed3ca60b048b4bf67cde2195961e0bba6f70bcbea9a2c4e133e34b46",
  "blockNumber": 3,
  "transactionIndex": 0,
  "from": "0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b",
  "to": "0x6295ee1b4f6dd65047762f924ecd367c17eabf8f",
  "value": BigNumber,
  "gas": 314159,
  "gasPrice": BigNumber,
  "input": "0x57cb2fc4"
}
*/

```

***

#### web3.eth.getTransactionReceipt

    web3.eth.getTransactionReceipt(hashString [, callback])

Returns the receipt of a transaction by transaction hash.

**Note** That the receipt is not available for pending transactions.


##### Parameters

1. `String` - The transaction hash.
2. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

##### Returns

`Object` - A transaction receipt object, or `null` when no receipt was found:

  - `blockHash`: `String`, 32 Bytes - hash of the block where this transaction was in.
  - `blockNumber`: `Number` - block number where this transaction was in.
  - `transactionHash`: `String`, 32 Bytes - hash of the transaction.
  - `transactionIndex`: `Number` - integer of the transactions index position in the block.
  - `from`: `String`, 20 Bytes - address of the sender.
  - `to`: `String`, 20 Bytes - address of the receiver. `null` when its a contract creation transaction.
  - `cumulativeGasUsed `: `Number ` - The total amount of gas used when this transaction was executed in the block.
  - `gasUsed `: `Number ` -  The amount of gas used by this specific transaction alone.
  - `contractAddress `: `String` - 20 Bytes - The contract address created, if the transaction was a contract creation, otherwise `null`.
  - `logs `:  `Array` - Array of log objects, which this transaction generated.

##### Example
```js
var receipt = web3.eth.getTransactionReceipt('0x9fc76417374aa880d4449a1f7f31ec597f00b1f6f3dd2d66f4c9c6c445836d8b');
console.log(receipt);
{
  "transactionHash": "0x9fc76417374aa880d4449a1f7f31ec597f00b1f6f3dd2d66f4c9c6c445836d8b",
  "transactionIndex": 0,
  "blockHash": "0xef95f2f1ed3ca60b048b4bf67cde2195961e0bba6f70bcbea9a2c4e133e34b46",
  "blockNumber": 3,
  "contractAddress": "0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b",
  "cumulativeGasUsed": 314159,
  "gasUsed": 30234,
  "logs": [{
         // logs as returned by getFilterLogs, etc.
     }, ...]
}
```

***

#### web3.eth.getTransactionCount

    web3.eth.getTransactionCount(addressHexString [, defaultBlock] [, callback])

Get the numbers of transactions sent from this address.

##### Parameters

1. `String` - The address to get the numbers of transactions from.
2. `Number|String` - (optional) If you pass this parameter it will not use the default block set with [web3.eth.defaultBlock](#web3ethdefaultblock).
3. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

##### Returns

`Number` - The number of transactions sent from the given address.

##### Example

```js
var number = web3.eth.getTransactionCount("0x407d73d8a49eeb85d32cf465507dd71d507100c1");
console.log(number); // 1
```

***

#### web3.eth.sendTransaction

    web3.eth.sendTransaction(transactionObject [, callback])

Sends a transaction to the network.

##### Parameters

1. `Object` - The transaction object to send:
  - `from`: `String` - The address for the sending account. Uses the [web3.eth.defaultAccount](#web3ethdefaultaccount) property, if not specified.
  - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation transaction.
  - `value`: `Number|String|BigNumber` - (optional) The value transferred for the transaction in Wei.
  - `gas`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The amount of gas to use for the transaction (unused gas is refunded).
  - `gasPrice`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The price of gas for this transaction in wei, defaults to the mean network gas price.
2. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

##### Returns

`String` - The 32 Bytes transaction hash as HEX string.

##### Example

```js

web3.eth.sendTransaction({from:web3.eth.coinbase,to:"0x7f9fade1c0d57a7af66ab4ead7c2eb7b11a91385",value:web3.toWei(1,"ether")}, function(err, hash) {
  if (!err)
    console.log(hash);
});
```

***

#### web3.eth.filter

```js
// can be 'latest' or 'pending'
var filter = web3.eth.filter(filterString);
// OR object are log filter options
var filter = web3.eth.filter(options);

// watch for changes
filter.watch(function(error, result){
  if (!error)
    console.log(result);
});

// Additionally you can start watching right away, by passing a callback:
web3.eth.filter(options, function(error, result){
  if (!error)
    console.log(result);
});
```

##### Parameters

1. `String|Object` - The string `"latest"` or `"pending"` to watch for changes in the latest block or pending transactions respectively. Or a filter options object as follows:
  * `fromBlock`: `Number|String` - The number of the earliest block (`latest` may be given to mean the most recent and `pending` currently mining, block). By default `latest`.
  * `toBlock`: `Number|String` - The number of the latest block (`latest` may be given to mean the most recent and `pending` currently mining, block). By default `latest`.
  * `address`: `String` - An address or a list of addresses to only get logs from particular account(s).
  * `topics`: `Array of Strings` - An array of values which must each appear in the log entries. The order is important, if you want to leave topics out use `null`, e.g. `[null, '0x00...']`. You can also pass another array for each topic with options for that topic e.g. `[null, ['option1', 'option2']]`

##### Returns

`Object` - A filter object with the following methods:

  * `filter.get(callback)`: Returns all of the log entries that fit the filter.
  * `filter.watch(callback)`: Watches for state changes that fit the filter and calls the callback. See [this note](#using-callbacks) for details.
  * `filter.stopWatching()`: Stops the watch and uninstalls the filter in the node. Should always be called once it is done.

##### Watch callback return value

- `String` - When using the `"latest"` parameter, it returns the block hash of the last incoming block.
- `String` - When using the `"pending"` parameter, it returns a transaction hash of the last add pending transaction.
- `Object` - When using manual filter options, it returns a log object as follows:
    - `logIndex`: `Number` - integer of the log index position in the block. `null` when its pending log.
    - `transactionIndex`: `Number` - integer of the transactions index position log was created from. `null` when its pending log.
    - `transactionHash`: `String`, 32 Bytes - hash of the transactions this log was created from. `null` when its pending log.
    - `blockHash`: `String`, 32 Bytes - hash of the block where this log was in. `null` when its pending. `null` when its pending log.
    - `blockNumber`: `Number` - the block number where this log was in. `null` when its pending. `null` when its pending log.
    - `address`: `String`, 32 Bytes - address from which this log originated.
    - `data`: `String` - contains one or more 32 Bytes non-indexed arguments of the log.
    - `topics`: `Array of Strings` - Array of 0 to 4 32 Bytes `DATA` of indexed log arguments. (In *solidity*: The first topic is the *hash* of the signature of the event (e.g. `Deposit(address,bytes32,uint256)`), except you declared the event with the `anonymous` specifier.)

**Note** For event filter return values see [Contract Events](#contract-events)

##### Example

```js
var filter = web3.eth.filter('pending');

filter.watch(function (error, log) {
  console.log(log); //  {"address":"0x0000000000000000000000000000000000000000", "data":"0x0000000000000000000000000000000000000000000000000000000000000000", ...}
});

// get all past logs again.
var myResults = filter.get(function(error, logs){ ... });

...

// stops and uninstalls the filter
filter.stopWatching();

```

***

#### web3.personal.unlockAccount

    web3.personal.unlockAccount
    // or async
    web3.personal.unlockAccount(web3.eth.accounts[1],"password",callback(error, result){ ... })

This property is used to unlock the account, sending transaction requires the sending account to be unlocked.

##### Parameters

1. `String` - The address of the account to unlock.
2. `String` - The pass phrase for that account.
3. `Number` - The duration for which to unlock the account in seconds.

##### Returns

`Boolean` - true | false , depending upon if the passphrase provided to unlock account is correct.

##### Example

```js
personal.unlockAccount(eth.coinbase,"password",function(err,accountUnlocked){
console.log("account unlocking status", accountUnlocked);  
})

```

***

#### web3.personal.newAccount

    web3.personal.newAccount

This property is used to create a new account in the local node directory.

##### Parameters

1. `String` - Password of the newly created account , used to generate the private key for the newly created account.

##### Returns

`String` - Account address of the newly created account.

##### Example

```js
personal.newAccount("myPassword");
"0x634a9702a76145ce5227e1cb90d434be7596092e"

```

***

#### web3.admin.addPeer

    web3.admin.addPeer(enodeString)

This property is used to add a remote node as a peer to your node.

##### Parameters

1. `String` - The enode string of the remote node.

##### Returns

`Boolean` - True if the local peer is able to add the remote node as peer , else false.

##### Example

```js
admin.addPeer("enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303")
returns true;

```

***
