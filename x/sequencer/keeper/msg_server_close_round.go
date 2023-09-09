package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"ibc_sequencer/x/sequencer/types"
)

func (k msgServer) CloseRound(goCtx context.Context, msg *types.MsgCloseRound) (*types.MsgCloseRoundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCloseRoundResponse{}, nil
}
