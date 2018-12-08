package cmd

import (
//	"fmt"
	//"strconv"
//	"github.com/eosforce/goeosforce/ecc"
//	"encoding/json"
	"github.com/eosforce/goeosforce/system"
	eos "github.com/eosforce/goeosforce"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
//	"strings"
)

var systemPushActionCmd = &cobra.Command{
	Use:   "pushaction [account name] [action name] [data]",
	Short: "Register an account as a block producer candidate.",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		

		accountName := toAccount(args[0], "account name")
		action_name := toActionName(args[1], "action name")
		
		//data
		api := getAPI()
		abi, err := api.GetABI(accountName)
		errorCheck("get ABI", err)

		data, _ := abi.ABI.EncodeAction(action_name, []byte(args[2]))
		
		requested, _ := permissionsToPermissionLevels(viper.GetStringSlice("push-action-cmd-requested-permissions"))
		errorCheck("requested permissions", err)
		if len(requested) == 0 {
			requested = []eos.PermissionLevel{
				{Actor: accountName, Permission: eos.PermissionName("active")},
			}
		}
		
		pushEOSCActions(api,
			system.NewAction(accountName, action_name,requested, eos.HexBytes(data).String()),
		)
	},
}

func init() {
	systemCmd.AddCommand(systemPushActionCmd)

	systemPushActionCmd.Flags().StringSliceP("requested-permissions", "", []string{}, "Permissions requested, specify multiple times or separated by a comma.")

	for _, flag := range []string{"requested-permissions"} {
		if err := viper.BindPFlag("push-action-cmd-"+flag, msigProposeCmd.Flags().Lookup(flag)); err != nil {
			panic(err)
		}
	}
}
