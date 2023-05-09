# Denomchain
Denomchain based on [Celestia Rollkit](https://rollkit.dev/docs/tutorials/building-and-deploying-a-rollup/) is created to register IBC denoms using the original information. It was created because a lot of developers had difficulties recognizing IBC token denoms in the Cosmos ecosystem.

## Install

```bash
# clone
git clone https://github.com/Jeongseup/denomchain.git && cd denomchain && git checkout v0.1.0 && make install

# check version
denomchaind version
```

## Start
```bash
# pre-requirement, running celestia bridge node
celestia light start --core.ip https://rpc-blockspacerace.pops.one --gateway --gateway.addr 127.0.0.1 --gateway.port 26659 --p2p.network blockspacerace


# start running denomchain node
NAMESPACE_ID=$(openssl rand -hex 8)
DA_BLOCK_HEIGHT=$(curl http://localhost:26657/block | jq -r '.result.block.header.height')

denomchaind start \
	--rollkit.aggregator true \
	--rollkit.da_layer celestia \
	--rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' \
	--rollkit.namespace_id $NAMESPACE_ID \
	--rollkit.da_start_height $DA_BLOCK_HEIGHT
```

## Available commands

```
--  Transaction --
Usage:
  denomchaind tx denomservice [flags]
  denomchaind tx denomservice [command]

Available Commands:
  set-denom   Broadcast message setDenom


-- Query --
Usage:
  denomchaind query denomservice [flags]
  denomchaind query denomservice [command]

Available Commands:
  list-denom  list all denom
  params      shows the parameters of the module
  show-denom  shows a denom
```


## Register a IBC Denom
```bash
# index: (chain-id/ibc/hex-codes)
# Usage: denomchaind tx denomservice set-denom [index] [port] [channel] [origin-denom]

# Example
denomchaind tx denomservice set-denom "cosmoshub-4/ibc/14f9bc3e44b8a9c1be1fb08980fab87034c9905ef17cf2f5008fc085218811cc" transfer channel-141 uosmo
```

## Query Registered IBC Denoms
```bash
# index: (chain-id/ibc/hex-codes)
# Usage: denomchaind query denomservice show-denom [index]

# Example
denomchaind q denomservice show-denom "cosmoshub-4/ibc/14f9bc3e44b8a9c1be1fb08980fab87034c9905ef17cf2f5008fc085218811cc"

'''  
>>> Expected Result
denom:
  channel: channel-141
  index: cosmoshub-4/ibc/14f9bc3e44b8a9c1be1fb08980fab87034c9905ef17cf2f5008fc085218811cc
  originDenom: uosmo
  port: transfer
'''


# Usage: denomchaind query denomservice list-denom
denomchaind q denomservice  list-denom
```