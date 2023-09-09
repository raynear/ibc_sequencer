package sequencer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ibc_sequencer/x/sequencer/keeper"
	"ibc_sequencer/x/sequencer/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the txPool
	for _, elem := range genState.TxPoolList {
		k.SetTxPool(ctx, elem)
	}
	// Set all the block
	for _, elem := range genState.BlockList {
		k.SetBlock(ctx, elem)
	}

	// Set block count
	k.SetBlockCount(ctx, genState.BlockCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.TxPoolList = k.GetAllTxPool(ctx)
	genesis.BlockList = k.GetAllBlock(ctx)
	genesis.BlockCount = k.GetBlockCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
