package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ibc_sequencer/x/sequencer/types"
)

// GetBlockCount get the total number of block
func (k Keeper) GetBlockCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BlockCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetBlockCount set the total number of block
func (k Keeper) SetBlockCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BlockCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBlock appends a block in the store with a new id and update the count
func (k Keeper) AppendBlock(
	ctx sdk.Context,
	block types.Block,
) uint64 {
	// Create the block
	count := k.GetBlockCount(ctx)

	// Set the ID of the appended value
	block.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))
	appendedValue := k.cdc.MustMarshal(&block)
	store.Set(GetBlockIDBytes(block.Id), appendedValue)

	// Update block count
	k.SetBlockCount(ctx, count+1)

	return count
}

// SetBlock set a specific block in the store
func (k Keeper) SetBlock(ctx sdk.Context, block types.Block) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))
	b := k.cdc.MustMarshal(&block)
	store.Set(GetBlockIDBytes(block.Id), b)
}

// GetBlock returns a block from its id
func (k Keeper) GetBlock(ctx sdk.Context, id uint64) (val types.Block, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))
	b := store.Get(GetBlockIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBlock removes a block from the store
func (k Keeper) RemoveBlock(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))
	store.Delete(GetBlockIDBytes(id))
}

// GetAllBlock returns all block
func (k Keeper) GetAllBlock(ctx sdk.Context) (list []types.Block) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Block
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBlockIDBytes returns the byte representation of the ID
func GetBlockIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBlockIDFromBytes returns ID in uint64 format from a byte array
func GetBlockIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
