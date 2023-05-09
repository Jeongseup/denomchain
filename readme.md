# denomchain


## Install

```
make install
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
$ denomchaind tx denomservice set-denom "cosmoshub-4/ibc/14f9bc3e44b8a9c1be1fb08980fab87034c9905ef17cf2f5008fc085218811cc" transfer channel-141 uosmo
```

## Query Registered IBC Denoms
```bash
# index: (chain-id/ibc/hex-codes)
# Usage: denomchaind query denomservice show-denom [index]

# Example
$ denomchaind q denomservice show-denom "cosmoshub-4/ibc/14f9bc3e44b8a9c1be1fb08980fab87034c9905ef17cf2f5008fc085218811cc"

'''  
>>> Expected Result
denom:
  channel: channel-141
  index: cosmoshub-4/ibc/14f9bc3e44b8a9c1be1fb08980fab87034c9905ef17cf2f5008fc085218811cc
  originDenom: uosmo
  port: transfer
'''


# Usage: denomchaind query denomservice list-denom
$ denomchaind q denomservice  list-denom
```