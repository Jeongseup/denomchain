package denomservice_test

import (
	"testing"

	keepertest "github.com/Jeongseup/denomchain/testutil/keeper"
	"github.com/Jeongseup/denomchain/testutil/nullify"
	"github.com/Jeongseup/denomchain/x/denomservice"
	"github.com/Jeongseup/denomchain/x/denomservice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DenomList: []types.Denom{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DenomserviceKeeper(t)
	denomservice.InitGenesis(ctx, *k, genesisState)
	got := denomservice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DenomList, got.DenomList)
	// this line is used by starport scaffolding # genesis/test/assert
}
