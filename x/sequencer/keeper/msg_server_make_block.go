package keeper

import (
	"context"

	"ibc_sequencer/x/sequencer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MakeBlock(goCtx context.Context, msg *types.MsgMakeBlock) (*types.MsgMakeBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	txs := k.GetAllTxPool(ctx)

	var block types.Block
	for i := 0; i < len(txs); i++ {
		if txs[i].Round == msg.Round && txs[i].Hash == "done" {
			block.Txs = append(block.Txs, txs[i].Payload)
		}
	}

	k.SetBlock(ctx, block)

	return &types.MsgMakeBlockResponse{}, nil
}
