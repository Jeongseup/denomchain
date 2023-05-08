#NAMESPACE_ID=$(openssl rand -hex 8)
NAMESPACE_ID=6232c6d502c1ebfd
echo $NAMESPACE_ID

DA_BLOCK_HEIGHT=$(curl http://localhost:26657/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

denomchaind start \
	--rollkit.aggregator true \
	--rollkit.da_layer celestia \
	--rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' \
	--rollkit.namespace_id $NAMESPACE_ID \
	--rollkit.da_start_height $DA_BLOCK_HEIGHT
