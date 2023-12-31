package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ibc_sequencer/testutil/keeper"
	"ibc_sequencer/x/sequencer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SequencerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
