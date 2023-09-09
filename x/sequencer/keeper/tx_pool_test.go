package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "ibc_sequencer/testutil/keeper"
	"ibc_sequencer/testutil/nullify"
	"ibc_sequencer/x/sequencer/keeper"
	"ibc_sequencer/x/sequencer/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTxPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TxPool {
	items := make([]types.TxPool, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTxPool(ctx, items[i])
	}
	return items
}

func TestTxPoolGet(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNTxPool(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTxPool(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTxPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNTxPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTxPool(ctx,
			item.Index,
		)
		_, found := keeper.GetTxPool(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTxPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNTxPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTxPool(ctx)),
	)
}
