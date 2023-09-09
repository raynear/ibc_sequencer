package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "ibc_sequencer/testutil/keeper"
	"ibc_sequencer/testutil/nullify"
	"ibc_sequencer/x/sequencer/keeper"
	"ibc_sequencer/x/sequencer/types"
)

func createNBlock(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Block {
	items := make([]types.Block, n)
	for i := range items {
		items[i].Id = keeper.AppendBlock(ctx, items[i])
	}
	return items
}

func TestBlockGet(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNBlock(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetBlock(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBlockRemove(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNBlock(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBlock(ctx, item.Id)
		_, found := keeper.GetBlock(ctx, item.Id)
		require.False(t, found)
	}
}

func TestBlockGetAll(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNBlock(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBlock(ctx)),
	)
}

func TestBlockCount(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNBlock(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetBlockCount(ctx))
}
