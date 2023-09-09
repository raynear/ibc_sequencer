package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	"ibc_sequencer/x/sequencer/types"
)

func (k msgServer) SendTlp(goCtx context.Context, msg *types.MsgSendTlp) (*types.MsgSendTlpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.TlpPacketData

	packet.Hash = msg.Hash
	packet.Tlp = msg.Tlp

	// Transmit the packet
	_, err := k.TransmitTlpPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendTlpResponse{}, nil
}
