package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"ibc_sequencer/x/sequencer/types"
)

func (k msgServer) MakeBlock(goCtx context.Context, msg *types.MsgMakeBlock) (*types.MsgMakeBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMakeBlockResponse{}, nil
}
