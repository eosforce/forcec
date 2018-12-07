package cmd

import (
	"fmt"

	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/system"
	"github.com/spf13/cobra"
)

var voteProducersCmd = &cobra.Command{
	Use:   "voteproducer [voter name] [bp name] [amount]",
	Short: "Cast your vote for 1 to 30 producers. View them with 'list-producers'.",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		voterName := toAccount(args[0], "voter name")

		bpName := toAccount(args[1], "bp name")
		quantity, _ := eos.NewEOSAssetFromString(args[2])
	
		api := getAPI()

		fmt.Printf("Voter [%s] voting for: %s\n", voterName, bpName)
		pushEOSCActions(api,
			system.NewVoteProducer(
				voterName,
				bpName,
				quantity,
			),
		)
	},
}

func init() {
	voteCmd.AddCommand(voteProducersCmd)
}
