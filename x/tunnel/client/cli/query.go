package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/chain/v2/x/tunnel/types"
)

const flagTunnelIsActiveFilter = "is-active"

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the tunnel module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetQueryTunnel(),
		GetQueryTunnels(),
		GetQueryCmdParams(),
	)

	return queryCmd
}

func GetQueryTunnel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tunnel [tunnel-id]",
		Short: "Query the tunnel by tunnel id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			tunnelID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.Tunnel(context.Background(), &types.QueryTunnelRequest{
				TunnelId: tunnelID,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQueryTunnels implements the query tunnels command.
func GetQueryTunnels() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tunnels",
		Short: "Query all tunnels",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			isActiveFilterFlag, err := cmd.Flags().GetBool(flagTunnelIsActiveFilter)
			if err != nil {
				return err
			}

			res, err := queryClient.Tunnels(context.Background(), &types.QueryTunnelsRequest{
				IsActive: isActiveFilterFlag,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().Bool(flagTunnelIsActiveFilter, false, "Filter tunnels by active status")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQueryCmdParams implements the query params command.
func GetQueryCmdParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
