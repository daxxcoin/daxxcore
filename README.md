## Daxxcoin Go

Official golang implementation of the Daxxcoin protocol.

[![API Reference](
)](https://godoc.org/github.com/daxxcoin/daxxcore)


## Building the source

Building geth requires both a Go and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

    make geth

or, to build the full suite of utilities:

    make all

## Executables

The daxxcore project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`geth`** | Our main Daxxcoin CLI client. It is the entry point into the Daxxcoin network, capable of running as a full node. It can be used by other processes as a gateway into the Daxxcoin network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. |
