`forcec` command-line swiss-army-knife
-------------------------------

forcec is fork from https://github.com/eoscanada/eosc

`forcec` is a cross-platform (Windows, Mac and Linux) command-line tool
for interacting with an EOSForce blockchain.

It contains tools for voting and a Vault to securely store private
keys.

It is based on the `goeosforce` library.

This first release holds simple tools, but a whole `cleos`-like
swiss-army-knife is being developed and will be released shortly after
mainnet launch.  Source code for most operations is already available
in this repository.


Installation
------------

1. Install from https://github.com/eosforce/forcec/releases

**or**

2. Build from source with:

```bash
go get -u -v github.com/eosforce/forcec/forcec
```