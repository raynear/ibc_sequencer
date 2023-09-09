package sequencer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ibc_sequencer/testutil/keeper"
	"ibc_sequencer/testutil/nullify"
	"ibc_sequencer/x/sequencer"
	"ibc_sequencer/x/sequencer/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		TxPoolList: []types.TxPool{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SequencerKeeper(t)
	sequencer.InitGenesis(ctx, *k, genesisState)
	got := sequencer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.TxPoolList, got.TxPoolList)
	// this line is used by starport scaffolding # genesis/test/assert
}
