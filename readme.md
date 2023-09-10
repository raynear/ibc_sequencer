# ibcsequencer

**ibcsequencer** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

First of all we run sequencer(mars) and rollup(venus) and relayer

```
// run two node. mars is sequencer. venus is rollup
$ ignite chain serve -c mars.yml -v
$ ignite chain serve -c venus.yml -v

// config IBC relayer
$ ignite relayer configure -a \
  --source-rpc "http://0.0.0.0:26657" \
  --source-faucet "http://0.0.0.0:4500" \
  --source-port "sequencer" \
  --source-version "sequencer-1" \
  --source-gasprice "0.0000025stake" \
  --source-prefix "cosmos" \
  --source-gaslimit 300000 \
  --target-rpc "http://0.0.0.0:26659" \
  --target-faucet "http://0.0.0.0:4501" \
  --target-port "sequencer" \
  --target-version "sequencer-1" \
  --target-gasprice "0.0000025stake" \
  --target-prefix "cosmos" \
  --target-gaslimit 300000

// activate relayer connection with nodes.
$ ignite relayer connect
```

## Running Sequencer example

### create transaction & send encrypted tx to sequencer(mars)

```
$ ibc_sequencerd tx sequencer create-tx-pool <Index> <hash<enc<Tx>>> <enc<Tx>> <Round> --from alice --chain-id mars --home ~/.mars
```

### stop receiving encrypted tx and send tx list commitment to rollup(venus) via IBC

```
$ ibc_sequencerd tx sequencer close-round <Round> --from alice --chain-id mars --home ~/.mars -b block
```

### send payload to rollup when rollup accepts sequencer's commitment

```
$ ibc_sequencerd tx sequencer send-payload sequencer channel-0 <Round> <hash<enc<Tx>[]>> --from alice --chain-id mars --home ~/.mars -y
```

### user can send time-lock puzzle via IBC so that rollup(venus) can decrypt the tx

```
$ ibc_sequencerd tx sequencer send-tlp sequencer channel-0 <hash<enc<Tx>>> <tlp<key>> --from alice --chain-id mars --home ~/.mars -y
```

### make the block with the decrypted tx

```
$ ibc_sequencerd tx sequencer make-block <Round> --from alice --chain-id venus --home ~/.venus -y
```

## Time-lock puzzle and encryption related

./tlp/createTLP.go generates time-lock puzzle with given seed phrase and minimum amount of seconds needed to be resolved and returns number of iterations and encryption key
./tlp/encryptTx.go encrypts the tx with given tx key and returns encrypted tx
./tlp/solveTLP.go receives seed phrase and number of iterations needed to obtain decryption key (===encyption key)
./decryptTx.go recieves encrypted tx and decryption key and decrypts the tx
