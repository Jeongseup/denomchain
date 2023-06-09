package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/Jeongseup/denomchain/testutil/keeper"
	"github.com/Jeongseup/denomchain/testutil/nullify"
	"github.com/Jeongseup/denomchain/x/denomservice/keeper"
	"github.com/Jeongseup/denomchain/x/denomservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDenom(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Denom {
	items := make([]types.Denom, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNewDenom(ctx, items[i])
	}
	return items
}

func TestDenomGet(t *testing.T) {
	keeper, ctx := keepertest.DenomserviceKeeper(t)
	items := createNDenom(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDenom(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestDenomGetAll(t *testing.T) {
	keeper, ctx := keepertest.DenomserviceKeeper(t)
	items := createNDenom(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDenom(ctx)),
	)
}
