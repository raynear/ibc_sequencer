package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ibc_sequencer/x/sequencer/types"
)

func (k Keeper) TxPoolAll(goCtx context.Context, req *types.QueryAllTxPoolRequest) (*types.QueryAllTxPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var txPools []types.TxPool
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	txPoolStore := prefix.NewStore(store, types.KeyPrefix(types.TxPoolKeyPrefix))

	pageRes, err := query.Paginate(txPoolStore, req.Pagination, func(key []byte, value []byte) error {
		var txPool types.TxPool
		if err := k.cdc.Unmarshal(value, &txPool); err != nil {
			return err
		}

		txPools = append(txPools, txPool)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTxPoolResponse{TxPool: txPools, Pagination: pageRes}, nil
}

func (k Keeper) TxPool(goCtx context.Context, req *types.QueryGetTxPoolRequest) (*types.QueryGetTxPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTxPool(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTxPoolResponse{TxPool: val}, nil
}
