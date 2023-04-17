package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/chain/v2/x/tss/types"
)

// NewTxCmd returns a root CLI command handler for all x/tss transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "TSS transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(MsgCreateGroupCmd())

	return txCmd
}

func MsgCreateGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-group [member1,member2,...] [threshold]",
		Args:  cobra.ExactArgs(2),
		Short: "Make a new group for tss module",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Make a new group for sign tx in tss module.
Example:
$ %s tx tss create-group band15mxunzureevrg646khnunhrl6nxvrj3eree5tz,band1p2t43jx3rz84y4z05xk8dcjjhzzeqnfrt9ua9v,band18f55l8hf4l7zvy8tx28n4r4nksz79p6lp4z305 2 --from mykey
`,
				version.AppName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			members := strings.Split(args[0], ",")

			threshold, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}

			msg := &types.MsgCreateGroup{
				Members:   members,
				Threshold: uint32(threshold),
				Sender:    clientCtx.GetFromAddress().String(),
			}
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
