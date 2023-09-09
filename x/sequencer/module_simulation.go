package sequencer

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"ibc_sequencer/testutil/sample"
	sequencersimulation "ibc_sequencer/x/sequencer/simulation"
	"ibc_sequencer/x/sequencer/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = sequencersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateTxPool = "op_weight_msg_tx_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTxPool int = 100

	opWeightMsgUpdateTxPool = "op_weight_msg_tx_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTxPool int = 100

	opWeightMsgDeleteTxPool = "op_weight_msg_tx_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTxPool int = 100

	opWeightMsgCloseRound = "op_weight_msg_close_round"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCloseRound int = 100

	opWeightMsgMakeBlock = "op_weight_msg_make_block"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMakeBlock int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	sequencerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		TxPoolList: []types.TxPool{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&sequencerGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTxPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTxPool, &weightMsgCreateTxPool, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTxPool = defaultWeightMsgCreateTxPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTxPool,
		sequencersimulation.SimulateMsgCreateTxPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTxPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTxPool, &weightMsgUpdateTxPool, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTxPool = defaultWeightMsgUpdateTxPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTxPool,
		sequencersimulation.SimulateMsgUpdateTxPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTxPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTxPool, &weightMsgDeleteTxPool, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTxPool = defaultWeightMsgDeleteTxPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTxPool,
		sequencersimulation.SimulateMsgDeleteTxPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCloseRound int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCloseRound, &weightMsgCloseRound, nil,
		func(_ *rand.Rand) {
			weightMsgCloseRound = defaultWeightMsgCloseRound
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCloseRound,
		sequencersimulation.SimulateMsgCloseRound(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMakeBlock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMakeBlock, &weightMsgMakeBlock, nil,
		func(_ *rand.Rand) {
			weightMsgMakeBlock = defaultWeightMsgMakeBlock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMakeBlock,
		sequencersimulation.SimulateMsgMakeBlock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
