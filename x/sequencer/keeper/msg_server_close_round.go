package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"

	"ibc_sequencer/x/sequencer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CloseRound(goCtx context.Context, msg *types.MsgCloseRound) (*types.MsgCloseRoundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	all := k.GetAllTxPool(ctx)

	hasher := sha256.New()

	for i := 0; i < len(all); i++ {
		hasher.Write([]byte(all[i].Payload))
	}

	hash := hasher.Sum(nil)

	var packet *types.MsgSendCommitment

	packet.Hash = bytes.NewBuffer(hash).String()
	packet.Round = msg.Round
	packet.Port = msg.Port
	packet.ChannelID = msg.ChannelID
	packet.TimeoutTimestamp = msg.TimeoutTimestamp

	_, err := k.SendCommitment(
		ctx,
		packet,
	)

	fmt.Println("error", err)

	return &types.MsgCloseRoundResponse{}, nil
}
