package cmd

import (

	//"strconv"
//	"github.com/eosforce/goeosforce/ecc"
	"encoding/json"
	"github.com/eosforce/goeosforce/system"
	eos "github.com/eosforce/goeosforce"
	"github.com/spf13/cobra"

)

var systemSetAccPermissionCmd = &cobra.Command{
	Use:   "setAccPermission [account name] [permission] [authority] [parent]",
	Short: "Register an account as a block producer candidate.",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		accountName := toAccount(args[0], "account name")
		permissionName := eos.PermissionName(args[1])
		
		var auth eos.Authority
		err := json.Unmarshal([]byte(args[2]), &auth)
		errorCheck("trans to auth", err)

		parent := eos.PermissionName(args[3])
		//
		action,_ := system.NewSetPermission(accountName,auth , permissionName, parent)
		api := getAPI()
		pushEOSCActions(api,
			action,
		)
	},
}

func init() {
	systemCmd.AddCommand(systemSetAccPermissionCmd)
}
