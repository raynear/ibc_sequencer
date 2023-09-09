package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendPayload = "send_payload"

var _ sdk.Msg = &MsgSendPayload{}

func NewMsgSendPayload(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	round uint64,
	hash string,
	encryptedTxs []string,
) *MsgSendPayload {
	return &MsgSendPayload{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Round:            round,
		Hash:             hash,
		EncryptedTxs:     encryptedTxs,
	}
}

func (msg *MsgSendPayload) Route() string {
	return RouterKey
}

func (msg *MsgSendPayload) Type() string {
	return TypeMsgSendPayload
}

func (msg *MsgSendPayload) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendPayload) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendPayload) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
