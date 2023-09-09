package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTxPool = "create_tx_pool"
	TypeMsgUpdateTxPool = "update_tx_pool"
	TypeMsgDeleteTxPool = "delete_tx_pool"
)

var _ sdk.Msg = &MsgCreateTxPool{}

func NewMsgCreateTxPool(
	creator string,
	index string,
	hash string,
	payload string,
	round uint64,

) *MsgCreateTxPool {
	return &MsgCreateTxPool{
		Creator: creator,
		Index:   index,
		Hash:    hash,
		Payload: payload,
		Round:   round,
	}
}

func (msg *MsgCreateTxPool) Route() string {
	return RouterKey
}

func (msg *MsgCreateTxPool) Type() string {
	return TypeMsgCreateTxPool
}

func (msg *MsgCreateTxPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTxPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTxPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTxPool{}

func NewMsgUpdateTxPool(
	creator string,
	index string,
	hash string,
	payload string,
	round uint64,

) *MsgUpdateTxPool {
	return &MsgUpdateTxPool{
		Creator: creator,
		Index:   index,
		Hash:    hash,
		Payload: payload,
		Round:   round,
	}
}

func (msg *MsgUpdateTxPool) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTxPool) Type() string {
	return TypeMsgUpdateTxPool
}

func (msg *MsgUpdateTxPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTxPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTxPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTxPool{}

func NewMsgDeleteTxPool(
	creator string,
	index string,

) *MsgDeleteTxPool {
	return &MsgDeleteTxPool{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteTxPool) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTxPool) Type() string {
	return TypeMsgDeleteTxPool
}

func (msg *MsgDeleteTxPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTxPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTxPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
