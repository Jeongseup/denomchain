#!/bin/sh

VALIDATOR_NAME=validator
CHAIN_ID=denom-1
KEY_NAME=validator
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

NAMESPACE_ID=$(openssl rand -hex 8)
echo $NAMESPACE_ID
DA_BLOCK_HEIGHT=$(curl https://localhost:26657/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

NODE_DAEMON=denomchaind
ignite chain build

$NODE_DAEMON tendermint unsafe-reset-all
$NODE_DAEMON init $VALIDATOR_NAME --chain-id $CHAIN_ID

$NODE_DAEMON keys add $KEY_NAME --keyring-backend test
$NODE_DAEMON add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
$NODE_DAEMON gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
$NODE_DAEMON collect-gentxs
$NODE_DAEMON start --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT
