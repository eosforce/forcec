// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (
	"encoding/json"
	"fmt"


	"strconv"

	"github.com/eosforce/goeosforce"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var voteListProducersCmd = &cobra.Command{
	Use:   "listbps [voter name]",
	Short: "Retrieve the list of registered producers.",
	Args:  cobra.MinimumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {
		api := getAPI()
	
		response, err := api.GetTableRows(
			eos.GetTableRowsRequest{
				Scope: args[0],
				Code:  "eosio",
				Table: "votes",
				JSON:  true,
				Limit: 5000,
			},
		)
		errorCheck("get table rows", err)
	
		if viper.GetBool("vote-list-cmd-json") {
			data, err := json.MarshalIndent(response.Rows, "", "    ")
			errorCheck("json marshal", err)
	
			fmt.Println(string(data))
		} else {
			var producers voteproducers
			err := json.Unmarshal(response.Rows, &producers)
			errorCheck("json marshaling", err)
	
			fmt.Println("List of producers registered to receive votes:")
			for _, p := range producers {
				fmt.Println("bpname:", p["bpname"]," staked: ", p["staked"]," unstaking:",p["unstaking"])
			}
			fmt.Printf("Total of %d registered producers\n", len(producers))
	
		}
	},
}

type voteproducers []map[string]interface{}

func (p voteproducers) Len() int      { return len(p) }
func (p voteproducers) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p voteproducers) Less(i, j int) bool {
	iv, _ := strconv.ParseFloat(p[i]["total_votes"].(string), 64)
	jv, _ := strconv.ParseFloat(p[j]["total_votes"].(string), 64)
	return iv > jv
}



func init() {
	voteCmd.AddCommand(voteListProducersCmd)

	voteListProducersCmd.Flags().BoolP("json", "j", false, "return producers info in json")

	for _, flag := range []string{"json"} {
		if err := viper.BindPFlag("vote-list-cmd-"+flag, voteListProducersCmd.Flags().Lookup(flag)); err != nil {
			panic(err)
		}
	}
}
