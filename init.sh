#!/bin/sh

VALIDATOR_NAME=validator
CHAIN_ID=denom-1
KEY_NAME=validator
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

NODE_DAEMON=denomchaind
ignite chain build

$NODE_DAEMON tendermint unsafe-reset-all
$NODE_DAEMON init $VALIDATOR_NAME --chain-id $CHAIN_ID

$NODE_DAEMON keys add $KEY_NAME --keyring-backend test
$NODE_DAEMON add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
$NODE_DAEMON gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
$NODE_DAEMON collect-gentxs
