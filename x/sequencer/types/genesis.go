package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:     PortID,
		TxPoolList: []TxPool{},
		BlockList:  []Block{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in txPool
	txPoolIndexMap := make(map[string]struct{})

	for _, elem := range gs.TxPoolList {
		index := string(TxPoolKey(elem.Index))
		if _, ok := txPoolIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for txPool")
		}
		txPoolIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in block
	blockIdMap := make(map[uint64]bool)
	blockCount := gs.GetBlockCount()
	for _, elem := range gs.BlockList {
		if _, ok := blockIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for block")
		}
		if elem.Id >= blockCount {
			return fmt.Errorf("block id should be lower or equal than the last id")
		}
		blockIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
