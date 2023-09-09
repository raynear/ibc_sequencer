package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"ibc_sequencer/x/sequencer/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				TxPoolList: []types.TxPool{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				BlockList: []types.Block{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				BlockCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated txPool",
			genState: &types.GenesisState{
				TxPoolList: []types.TxPool{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated block",
			genState: &types.GenesisState{
				BlockList: []types.Block{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid block count",
			genState: &types.GenesisState{
				BlockList: []types.Block{
					{
						Id: 1,
					},
				},
				BlockCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
