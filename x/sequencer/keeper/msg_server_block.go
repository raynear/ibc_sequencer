package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"ibc_sequencer/x/sequencer/types"
)

func (k msgServer) CreateBlock(goCtx context.Context, msg *types.MsgCreateBlock) (*types.MsgCreateBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var block = types.Block{
		Creator: msg.Creator,
		Txs:     msg.Txs,
	}

	id := k.AppendBlock(
		ctx,
		block,
	)

	return &types.MsgCreateBlockResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBlock(goCtx context.Context, msg *types.MsgUpdateBlock) (*types.MsgUpdateBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var block = types.Block{
		Creator: msg.Creator,
		Id:      msg.Id,
		Txs:     msg.Txs,
	}

	// Checks that the element exists
	val, found := k.GetBlock(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBlock(ctx, block)

	return &types.MsgUpdateBlockResponse{}, nil
}

func (k msgServer) DeleteBlock(goCtx context.Context, msg *types.MsgDeleteBlock) (*types.MsgDeleteBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetBlock(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBlock(ctx, msg.Id)

	return &types.MsgDeleteBlockResponse{}, nil
}
