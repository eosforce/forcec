package cmd

import (
	"fmt"

//	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/system"
	"github.com/spf13/cobra"
)

var voteUnfreezeCmd = &cobra.Command{
	Use:   "unfreeze [voter name] [bp name]",
	Short: "Receive dividends on the bp",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		voterName := toAccount(args[0], "voter name")

		bpName := toAccount(args[1], "bp name")
	
		api := getAPI()

		fmt.Printf("Voter [%s] voting for: %s\n", voterName, bpName)
		pushEOSCActions(api,
			system.NewUnfreeze(
				voterName,
				bpName,
			),
		)
	},
}

func init() {
	voteCmd.AddCommand(voteUnfreezeCmd)
}
