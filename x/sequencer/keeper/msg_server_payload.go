package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"

	"ibc_sequencer/x/sequencer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
)

func (k msgServer) SendPayload(goCtx context.Context, msg *types.MsgSendPayload) (*types.MsgSendPayloadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// reduce parameter
	all := k.GetAllTxPool(ctx)

	hasher := sha256.New()

	txs := []string{}
	for i := 0; i < len(all); i++ {
		txs = append(txs, all[i].Payload)
		hasher.Write([]byte(all[i].Payload))
	}
	// Construct the packet
	var packet types.PayloadPacketData

	hash := hasher.Sum(nil)

	packet.Round = msg.Round
	packet.Hash = bytes.NewBuffer(hash).String()
	packet.EncryptedTxs = txs

	// Transmit the packet
	_, err := k.TransmitPayloadPacket(
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

	return &types.MsgSendPayloadResponse{}, nil
}
