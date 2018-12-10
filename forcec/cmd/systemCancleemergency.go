package cmd

import (
//	"fmt"
//	"strconv"
	//"github.com/eosforce/goeosforce/ecc"
	"github.com/eosforce/goeosforce/system"
	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
)

var systemCancleemergencyCmd = &cobra.Command{
	Use:   "cancleemergency [bp_name]",
	Short: "Cancle the status of the chain is an emergency",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		accountName := toAccount(args[0], "bp name")

		pushEOSCActions(api,
			system.NewSetemergency(accountName, false),
		)
	},
}

func init() {
	systemCmd.AddCommand(systemCancleemergencyCmd)
}
