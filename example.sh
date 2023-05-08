denomchaind tx denomservice set-denom "cosmoshub-4/ibc/14f9bc3e44b8a9c1be1fb08980fab87034c9905ef17cf2f5008fc085218811cc" transfer channel-141 uosmo \
	--from validator \
       	--chain-id denom-1

 denomchaind q denomservice  list-denom
