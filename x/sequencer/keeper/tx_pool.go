package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ibc_sequencer/x/sequencer/types"
)

// SetTxPool set a specific txPool in the store from its index
func (k Keeper) SetTxPool(ctx sdk.Context, txPool types.TxPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxPoolKeyPrefix))
	b := k.cdc.MustMarshal(&txPool)
	store.Set(types.TxPoolKey(
		txPool.Index,
	), b)
}

// GetTxPool returns a txPool from its index
func (k Keeper) GetTxPool(
	ctx sdk.Context,
	index string,

) (val types.TxPool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxPoolKeyPrefix))

	b := store.Get(types.TxPoolKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTxPool removes a txPool from the store
func (k Keeper) RemoveTxPool(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxPoolKeyPrefix))
	store.Delete(types.TxPoolKey(
		index,
	))
}

// GetAllTxPool returns all txPool
func (k Keeper) GetAllTxPool(ctx sdk.Context) (list []types.TxPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxPoolKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TxPool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
