package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetDenom = "set_denom"

var _ sdk.Msg = &MsgSetDenom{}

func NewMsgSetDenom(creator string, index string, port string, channel string, originDenom string) *MsgSetDenom {
	return &MsgSetDenom{
		Creator:     creator,
		Index:       index,
		Port:        port,
		Channel:     channel,
		OriginDenom: originDenom,
	}
}

func (msg *MsgSetDenom) Route() string {
	return RouterKey
}

func (msg *MsgSetDenom) Type() string {
	return TypeMsgSetDenom
}

func (msg *MsgSetDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
