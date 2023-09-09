package keeper

import (
	"context"

	"ibc_sequencer/x/sequencer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MakeBlock(goCtx context.Context, msg *types.MsgMakeBlock) (*types.MsgMakeBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: raynear
	_ = ctx
	// make block from stored dec_tx
	// k.CreateBlock(ctx, )
	// 저장한거 걍 가져다가 block만들기 - DB에 저장하기

	return &types.MsgMakeBlockResponse{}, nil
}
