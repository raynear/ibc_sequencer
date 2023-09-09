package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"ibc_sequencer/x/sequencer/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateTxPool())
	cmd.AddCommand(CmdUpdateTxPool())
	cmd.AddCommand(CmdDeleteTxPool())
	cmd.AddCommand(CmdSendCommitment())
	cmd.AddCommand(CmdSendPayload())
	cmd.AddCommand(CmdSendTlp())
	cmd.AddCommand(CmdCloseRound())
	cmd.AddCommand(CmdMakeBlock())
	cmd.AddCommand(CmdCreateBlock())
	cmd.AddCommand(CmdUpdateBlock())
	cmd.AddCommand(CmdDeleteBlock())
	// this line is used by starport scaffolding # 1

	return cmd
}
