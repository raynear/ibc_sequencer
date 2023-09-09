package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ibc_sequencer/x/sequencer/types"
)

func (k Keeper) BlockAll(goCtx context.Context, req *types.QueryAllBlockRequest) (*types.QueryAllBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var blocks []types.Block
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	blockStore := prefix.NewStore(store, types.KeyPrefix(types.BlockKey))

	pageRes, err := query.Paginate(blockStore, req.Pagination, func(key []byte, value []byte) error {
		var block types.Block
		if err := k.cdc.Unmarshal(value, &block); err != nil {
			return err
		}

		blocks = append(blocks, block)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBlockResponse{Block: blocks, Pagination: pageRes}, nil
}

func (k Keeper) Block(goCtx context.Context, req *types.QueryGetBlockRequest) (*types.QueryGetBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	block, found := k.GetBlock(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetBlockResponse{Block: block}, nil
}
