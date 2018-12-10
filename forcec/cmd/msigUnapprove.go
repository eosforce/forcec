// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (
	"encoding/json"
	eos "github.com/eosforce/goeosforce"
	 "github.com/eosforce/goeosforce/msig"
	 "github.com/spf13/cobra"
)

// msigUnapproveCmd represents the `eosio.msig::unapprove` command
var msigUnapproveCmd = &cobra.Command{
	Use:   "unapprove [proposer] [proposal_name] [permissions]",
	Short: "Unapprove a transaction in the eosio.msig contract",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		proposer_name := toAccount(args[0], "proposer")
		proposal_name := toName(args[1], "proposal name")

		var trx_permissions eos.PermissionLevel
		errReq := json.Unmarshal([]byte(args[2]), &trx_permissions)
		errorCheck("trx_permissions", errReq)

		pushEOSCActions(api,
			msig.NewUnapprove(proposer_name, proposal_name, trx_permissions),
		)

	},
}

func init() {
	msigCmd.AddCommand(msigUnapproveCmd)
}
