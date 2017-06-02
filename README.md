## Daxxcoin Go

Official golang implementation of the Daxxcoin protocol.

[![API Reference](
)](https://godoc.org/github.com/daxxcoin/daxxcore)


## Building the source

Building geth requires both a Go and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

    make geth
    
## Executables

The daxxcore project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`geth`** | Our main Daxxcoin CLI client. It is the entry point into the Daxxcoin network, capable of running as a full node. It can be used by other processes as a gateway into the Daxxcoin network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. |

## Connecting to the DaxxCoin network

To connect to the Daxxcoin network , please follow the steps in order - 

1) Move into a new directory with <b>read+write access</b> , and copy into that directory both the files from the <b>daxxcore/networkFiles</b> directory.

2) In the new directory run the following to create the <b>Daxxcoin blockchain data folder and a keystore folder which will keep track of the accounts created on your node.</b>

    <b>./path-to-daxxcore/build/bin/geth --datadir myNode --networkid 11199 --port 30333 init CustomGenesis.json</b>

3) After you have created the geth and keystore folder , just move the <b>static-nodes.json file</b> into the <b>myNode directory.</b>

4) Now you can run the local node in <b>interactive mode</b> using the following command - 

    <b>./path-to-daxxcore/build/bin/geth --datadir myNode --networkid 11199 --port 30333 console</b> 

5) To make sure that you are connected with the <b>DaxxCoin Blockchain</b> , see that the node starts <b>synchronising with peers and starts downloading the blockchain data from peers.</b>

6) To connect with the local daxxcoin node , <b>using rpc</b> , use the following command after when you have successfully initiated a Daxxcoin node data directory as given in Step 2. Using rpc you are availed all the <b>methods listed in the javascriptAPI.md file in the daxxcore directory.</b>

    <b>./path-to-daxxcore/build/bin/geth --datadir myNode --networkid 11199 --port 30333 --rpc --rpcapi "admin,personal,web3,eth,net" --rpcaddr "127.0.0.1" --rpcport 8888 console</b>

7) If you encounter a <b>port conflict</b> , you can choose any other ports which are available for communication on your machine instead of the default paths given in the above commands respectively for <b>--port and --rpcport.</b>

8) <b>The keystore folder is where your account files are stored . So always keep a backup of this folder and also try to keep remember-able passphrases , because there is no method to get back the passphrase when you lose it.</b>
