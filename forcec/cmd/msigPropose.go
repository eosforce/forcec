// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (
	"encoding/json"
	//"errors"
	//"io/ioutil"

	"github.com/eosforce/goeosforce/system"
	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/msig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// msigProposeCmd represents the msigPropose command
var msigProposeCmd = &cobra.Command{
	Use:   "propose [proposal_name] [requested_permissions] [trx_permissions] [contract] [action] [data] [proposer]",
	Short: "Propose a new transaction in the eosio.msig contract",
	Long: `Propose a new transaction in the eosio.msig contract

Pass --requested-permissions
`,
	Args: cobra.ExactArgs(7),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		proposal_name := toName(args[0], "proposal name")
		
		var requested_permissions []eos.PermissionLevel
		err := json.Unmarshal([]byte(args[1]), &requested_permissions)
		errorCheck("requested_permissions", err)

		var trx_permissions []eos.PermissionLevel
		errReq := json.Unmarshal([]byte(args[2]), &trx_permissions)
		errorCheck("trx_permissions", errReq)

		accountName := toAccount(args[3], "contract")
		action_name := toActionName(args[4], "action")

		abi, err := api.GetABI(accountName)
		errorCheck("get ABI", err)

		data, _ := abi.ABI.EncodeAction(action_name, []byte(args[5]))
		action := system.NewAction(accountName, action_name,trx_permissions, eos.HexBytes(data).String())

		proposer_name := toAccount(args[6], "proposer")

		trx := getEOSCTransaction(api,action)

		pushEOSCActions(api,
			msig.NewPropose(proposer_name, proposal_name, requested_permissions, trx),
		)
	},
}

func init() {
	msigCmd.AddCommand(msigProposeCmd)

	msigProposeCmd.Flags().StringSliceP("requested-permissions", "", []string{}, "Permissions requested, specify multiple times or separated by a comma.")

	for _, flag := range []string{"requested-permissions"} {
		if err := viper.BindPFlag("msig-propose-cmd-"+flag, msigProposeCmd.Flags().Lookup(flag)); err != nil {
			panic(err)
		}
	}

}
