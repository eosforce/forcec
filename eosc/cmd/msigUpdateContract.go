// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (
	"fmt"
	"sort"
	"strings"

	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/msig"
	"github.com/spf13/cobra"
	system "github.com/eosforce/goeosforce/system"
)

// msigProposeCmd represents the msigPropose command
var msigUpdatSysCmd = &cobra.Command{
	Use:   "updatesys [proposer] [proposal name] [system account] [wasmPath] [abiPath]",
	Short: "Propose a new transaction in the eosio.msig contract",
	Long: `Propose a new transaction in the eosio.msig contract  `,
	Args: cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		proposer := toAccount(args[0], "proposer")
		proposalName := toName(args[1], "proposal name")

		sysaccount := toAccount(args[2], "system account")
		wasmPath := args[3]
		abiPath := args[4]

		
		tx,_ := system.NewSetCodeTx(sysaccount,wasmPath,abiPath);
		fee, err := GetFeeByTrx(tx)
		if err != nil {
			fmt.Println("Error get fee:", err)
			fee = eos.NewEOSAsset(15000)
		}
	
		tx.Fee = fee
	
		var requested []eos.PermissionLevel
		out, err := requestProducers(api)
		errorCheck("recursing to get producers accounts", err)

		for el := range out {
			chunks := strings.Split(el, "@")
			requested = append(requested, eos.PermissionLevel{
				Actor:      eos.AccountName(chunks[0]),
				Permission: eos.PermissionName(chunks[1]),
			})
			fmt.Println(chunks[0],"---el---",chunks[1])
		}

		sort.Slice(requested, func(i, j int) bool {
			el1 := requested[i]
			el2 := requested[j]
			if el1.Actor < el2.Actor {
				return true
			}
			if el1.Actor > el2.Actor {
				return false
			}
			return el1.Permission < el2.Permission
		})

		pushEOSCActions(api,
			msig.NewPropose(proposer, proposalName, requested, tx),
		)
	},
}

func init() {
	msigCmd.AddCommand(msigUpdatSysCmd)
}
