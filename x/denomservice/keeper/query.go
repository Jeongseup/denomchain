package keeper

import (
	"github.com/Jeongseup/denomchain/x/denomservice/types"
)

var _ types.QueryServer = Keeper{}
