package keeper_test

import (
	"testing"

	testkeeper "github.com/Jeongseup/denomchain/testutil/keeper"
	"github.com/Jeongseup/denomchain/x/denomservice/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DenomserviceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
