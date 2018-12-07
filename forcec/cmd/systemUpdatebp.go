package cmd

import (
	"fmt"
	"strconv"
	"github.com/eosforce/goeosforce/ecc"
	"github.com/eosforce/goeosforce/system"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var systemRegisterProducerCmd = &cobra.Command{
	Use:   "updatebp [account_name] [producer_key] [commission_rate] [website_url]",
	Short: "Register an account as a block producer candidate.",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		accountName := toAccount(args[0], "account name")
		publicKey, err := ecc.NewPublicKey(args[1])
		errorCheck(fmt.Sprintf("%q invalid public key", args[1]), err)
		commissionRata,_ :=  strconv.Atoi(args[2])
		websiteURL := args[3]

		pushEOSCActions(api,
			system.NewRegProducer(accountName, publicKey,websiteURL, int32(commissionRata)),
		)
	},
}

func init() {
	systemCmd.AddCommand(systemRegisterProducerCmd)

	systemRegisterProducerCmd.Flags().IntP("location", "", 0, "Location number (reserved)")

	for _, flag := range []string{"location"} {
		if err := viper.BindPFlag("system-regproducer-cmd-"+flag, systemRegisterProducerCmd.Flags().Lookup(flag)); err != nil {
			panic(err)
		}
	}

}
