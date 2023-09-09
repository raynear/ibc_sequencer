package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"ibc_sequencer/x/sequencer/keeper"
	"ibc_sequencer/x/sequencer/types"
)

func SimulateMsgMakeBlock(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMakeBlock{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MakeBlock simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MakeBlock simulation not implemented"), nil, nil
	}
}
