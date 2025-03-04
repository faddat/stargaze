package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/public-awesome/stargaze/x/stake/types"
	"github.com/spf13/cobra"
)

// NewStakeTxCmd returns the stake command
func NewStakeTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake [vendor-id] [post-id] [amount] [validator-address] --from [key]",
		Args:  cobra.MinimumNArgs(4),
		Short: "Stake on a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Stake on a post.
Example:
$ %s tx stake stake 1 "2" 500 --from mykey
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			delegator := clientCtx.GetFromAddress()

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			postID := args[1]
			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				panic("invalid amount")
			}

			valAddrStr := args[3]
			validator, err := sdk.ValAddressFromBech32(valAddrStr)
			if err != nil {
				return err
			}

			msg := types.NewMsgStake(
				uint32(vendorID), postID, delegator, validator, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewUnstakeTxCmd returns the unstake command
func NewUnstakeTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unstake [vendor-id] [post-id] [amount] --from [key]",
		Args:  cobra.MinimumNArgs(3),
		Short: "Unstake from a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Stake on a post.
Example:
$ %s tx stake unstake 1 "2" 500 --from mykey
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			delegator := clientCtx.GetFromAddress()

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			postID := args[1]
			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				panic("invalid amount")
			}

			msg := types.NewMsgUnstake(uint32(vendorID), postID, delegator, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
