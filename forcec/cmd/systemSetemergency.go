package cmd

import (
	//"fmt"
	//"strconv"
	//"github.com/eosforce/goeosforce/ecc"
	"github.com/eosforce/goeosforce/system"
	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
)

var systemSetemergencyCmd = &cobra.Command{
	Use:   "setemergency [bp_name]",
	Short: "Setting the status of the chain is an emergency",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		accountName := toAccount(args[0], "bp name")

		pushEOSCActions(api,
			system.NewSetemergency(accountName, true),
		)
	},
}

func init() {
	systemCmd.AddCommand(systemSetemergencyCmd)
}
