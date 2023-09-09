package keeper

import (
	"ibc_sequencer/x/sequencer/types"
)

var _ types.QueryServer = Keeper{}
