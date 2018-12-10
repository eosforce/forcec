// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (
	// "fmt"
	// "os"
	"encoding/json"
	eos "github.com/eosforce/goeosforce"
	 "github.com/eosforce/goeosforce/msig"
	 "github.com/spf13/cobra"
)

// msigApproveCmd represents the `eosio.msig::approve` command
var msigApproveCmd = &cobra.Command{
	Use:   "approve [proposer] [proposal_name] [permissions]",
	Short: "Approve a transaction in the eosio.msig contract",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		proposer_name := toAccount(args[0], "proposer")
		proposal_name := toName(args[1], "proposal name")

		var trx_permissions eos.PermissionLevel
		errReq := json.Unmarshal([]byte(args[2]), &trx_permissions)
		errorCheck("trx_permissions", errReq)

		pushEOSCActions(api,
			msig.NewApprove(proposer_name, proposal_name, trx_permissions),
		)
	},
}

func init() {
	msigCmd.AddCommand(msigApproveCmd)
}
