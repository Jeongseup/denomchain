package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Jeongseup/denomchain/x/denomservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetDenom(goCtx context.Context, msg *types.MsgSetDenom) (*types.MsgSetDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// chain_id/ibc_denom
	// testnet/IBC/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9
	slice := strings.Split(msg.Index, "/")
	// return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("here %d, %s, %s, %s", len(slice), slice[0], slice[1], slice[2]))

	switch len(slice) {
	case 0:
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("Incorrect Index, lengh is 0, check your message is %s", slice[0]))
	case 1:
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("Incorrect Index, lengh is 1, check your message is %s and %s", slice[0], slice[1]))
	case 2:
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("Incorrect Index, lengh is 1, check your message is %s and %s and %s", slice[0], slice[1], slice[2]))
	case 3:
		// chain-id
		chainID := slice[0]
		_ = chainID

		// ibc prefix
		ibcPrefix := strings.ToLower(slice[1])
		// ibc denom
		ibcDenom := slice[2]

		if "ibc" != ibcPrefix {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("Incorrect Index, check,ldaer your message is %s and %s", slice[0], slice[1]))
		}

		// (msg.Port, msg.Channel, msg.OriginDenom)
		data := msg.GetPort() + "/" + msg.GetChannel() + "/" + msg.GetOriginDenom()
		hash := sha256.New()
		hash.Write([]byte(data))
		md := hash.Sum(nil)
		mdStr := hex.EncodeToString(md)

		// If the denom index is not correct, throw an error
		if ibcDenom != mdStr {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("Incorrect Index, check, your ibc denom is %s\n your input string is %s, expected denom %s", slice[2], data, mdStr))
		}

		// check index chain-id/ibc/denom
		// Try getting denom information from the store
		denom, _ := k.GetDenom(ctx, msg.Index)
		if denom.Index != "" {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "already registered")
		}

		// Otherwise, create an updated whois record
		newDenom := types.Denom{
			Index:       msg.Index,
			Port:        msg.Port,
			Channel:     msg.Channel,
			OriginDenom: msg.OriginDenom,
		}

		// Write new denom to the store
		k.SetNewDenom(ctx, newDenom)
		return &types.MsgSetDenomResponse{}, nil

	default:
		// incorrect index
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("Incorrect Index, length is default, your message is %s", msg.Index))
	}
}
