// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (

	"github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/ecc"
	"github.com/eosforce/goeosforce/system"
	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
)

var systemNewAccountCmd = &cobra.Command{
	Use:   "newaccount [creator] [new_account_name] [owner_key] [active_key]",
	Short: "Create a new account.",
	Long: `Create a new account on the blockchain with initial resources`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		creator := toAccount(args[0], "creator")
		newAccount := toAccount(args[1], "new account name")

		var actions []*eos.Action
		owner_key,_ := ecc.NewPublicKey(args[2])
		active_key,_ := ecc.NewPublicKey(args[3])

		a := eos.Authority{
			Threshold: 1,
			Keys: []eos.KeyWeight{
				eos.KeyWeight{
					PublicKey: owner_key,
					Weight:    1,
				},
			},
		}

		b := eos.Authority{
			Threshold: 1,
			Keys: []eos.KeyWeight{
				eos.KeyWeight{
					PublicKey: active_key,
					Weight:    1,
				},
			},
		}

		actions = append(actions, system.NewCustomNewAccount(creator, newAccount, a,b))

		api := getAPI()
		pushEOSCActions(api, actions...)
		
	},

}

func init() {
	systemCmd.AddCommand(systemNewAccountCmd)



}
