package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"ibc_sequencer/x/sequencer/types"
)

func (k msgServer) CreateTxPool(goCtx context.Context, msg *types.MsgCreateTxPool) (*types.MsgCreateTxPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTxPool(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var txPool = types.TxPool{
		Creator: msg.Creator,
		Index:   msg.Index,
		Hash:    msg.Hash,
		Payload: msg.Payload,
		Round:   msg.Round,
	}

	k.SetTxPool(
		ctx,
		txPool,
	)
	return &types.MsgCreateTxPoolResponse{}, nil
}

func (k msgServer) UpdateTxPool(goCtx context.Context, msg *types.MsgUpdateTxPool) (*types.MsgUpdateTxPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTxPool(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var txPool = types.TxPool{
		Creator: msg.Creator,
		Index:   msg.Index,
		Hash:    msg.Hash,
		Payload: msg.Payload,
		Round:   msg.Round,
	}

	k.SetTxPool(ctx, txPool)

	return &types.MsgUpdateTxPoolResponse{}, nil
}

func (k msgServer) DeleteTxPool(goCtx context.Context, msg *types.MsgDeleteTxPool) (*types.MsgDeleteTxPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTxPool(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTxPool(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteTxPoolResponse{}, nil
}
