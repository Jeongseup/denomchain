package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Jeongseup/denomchain/testutil/keeper"
	"github.com/Jeongseup/denomchain/x/denomservice/keeper"
	"github.com/Jeongseup/denomchain/x/denomservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DenomserviceKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
