#!/bin/bash

KEY="dev0"
CHAINID="highbury_710-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t fury-datadir.XXXXX)

echo "create and add new keys"
./furyd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Fury with moniker=$MONIKER and chain-id=$CHAINID"
./furyd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./furyd add-genesis-account \
"$(./furyd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000afury,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./furyd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./furyd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./furyd validate-genesis --home $DATA_DIR

echo "starting fury node $i in background ..."
./furyd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started fury node"
tail -f /dev/null