package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMakeBlock = "make_block"

var _ sdk.Msg = &MsgMakeBlock{}

func NewMsgMakeBlock(creator string, round uint64) *MsgMakeBlock {
	return &MsgMakeBlock{
		Creator: creator,
		Round:   round,
	}
}

func (msg *MsgMakeBlock) Route() string {
	return RouterKey
}

func (msg *MsgMakeBlock) Type() string {
	return TypeMsgMakeBlock
}

func (msg *MsgMakeBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMakeBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMakeBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
