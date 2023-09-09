package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBlock = "create_block"
	TypeMsgUpdateBlock = "update_block"
	TypeMsgDeleteBlock = "delete_block"
)

var _ sdk.Msg = &MsgCreateBlock{}

func NewMsgCreateBlock(creator string, txs []string) *MsgCreateBlock {
	return &MsgCreateBlock{
		Creator: creator,
		Txs:     txs,
	}
}

func (msg *MsgCreateBlock) Route() string {
	return RouterKey
}

func (msg *MsgCreateBlock) Type() string {
	return TypeMsgCreateBlock
}

func (msg *MsgCreateBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBlock{}

func NewMsgUpdateBlock(creator string, id uint64, txs []string) *MsgUpdateBlock {
	return &MsgUpdateBlock{
		Id:      id,
		Creator: creator,
		Txs:     txs,
	}
}

func (msg *MsgUpdateBlock) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBlock) Type() string {
	return TypeMsgUpdateBlock
}

func (msg *MsgUpdateBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBlock{}

func NewMsgDeleteBlock(creator string, id uint64) *MsgDeleteBlock {
	return &MsgDeleteBlock{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteBlock) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBlock) Type() string {
	return TypeMsgDeleteBlock
}

func (msg *MsgDeleteBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
