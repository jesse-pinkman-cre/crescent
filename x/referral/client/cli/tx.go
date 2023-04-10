package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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

	cmd.AddCommand(
		GetAddReferralCmd(),
		GetSetReferralCmd(),
	)

	return cmd
}

func GetAddReferralCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-referral [code]",
		Args:  cobra.ExactArgs(1),
		Short: "Create my referral code",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			//TODO JIHON: input validation
			// coins, err := sdk.ParseCoinsNormalized(args[0])
			// if err != nil {
			// 	return fmt.Errorf("invalid coins: %w", err)
			// }
			code := args[0]
			msg := types.NewMsgAddReferral(clientCtx.GetFromAddress(), code)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetSetReferralCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-parent [parent]",
		Args:  cobra.ExactArgs(1),
		Short: "Set my parent code",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			//TODO JIHON: input validation
			// coins, err := sdk.ParseCoinsNormalized(args[0])
			// if err != nil {
			// 	return fmt.Errorf("invalid coins: %w", err)
			// }
			parent := args[0]
			msg := types.NewMsgSetReferral(clientCtx.GetFromAddress(), parent)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
