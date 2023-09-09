package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCloseRound = "close_round"

var _ sdk.Msg = &MsgCloseRound{}

func NewMsgCloseRound(creator string, round uint64) *MsgCloseRound {
	return &MsgCloseRound{
		Creator: creator,
		Round:   round,
	}
}

func (msg *MsgCloseRound) Route() string {
	return RouterKey
}

func (msg *MsgCloseRound) Type() string {
	return TypeMsgCloseRound
}

func (msg *MsgCloseRound) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCloseRound) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCloseRound) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
