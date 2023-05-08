package cli

import (
	"strconv"

	"github.com/Jeongseup/denomchain/x/denomservice/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-denom [index] [port] [channel] [origin-denom]",
		Short: "Broadcast message setDenom",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIndex := args[0]
			argPort := args[1]
			argChannel := args[2]
			argOriginDenom := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetDenom(
				clientCtx.GetFromAddress().String(),
				argIndex,
				argPort,
				argChannel,
				argOriginDenom,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
