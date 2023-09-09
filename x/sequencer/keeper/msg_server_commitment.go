package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	"ibc_sequencer/x/sequencer/types"
)

func (k msgServer) SendCommitment(goCtx context.Context, msg *types.MsgSendCommitment) (*types.MsgSendCommitmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.CommitmentPacketData

	packet.Round = msg.Round
	packet.Hash = msg.Hash

	// Transmit the packet
	_, err := k.TransmitCommitmentPacket(
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

	return &types.MsgSendCommitmentResponse{}, nil
}
