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

var systemListProducersCmd = &cobra.Command{
	Use:   "listbps",
	Short: "Retrieve the list of registered producers.",
	Run:   run,
}

type producers []map[string]interface{}

func (p producers) Len() int      { return len(p) }
func (p producers) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p producers) Less(i, j int) bool {
	iv, _ := strconv.ParseFloat(p[i]["total_votes"].(string), 64)
	jv, _ := strconv.ParseFloat(p[j]["total_votes"].(string), 64)
	return iv > jv
}

var run = func(cmd *cobra.Command, args []string) {
	api := getAPI()

	response, err := api.GetTableRows(
		eos.GetTableRowsRequest{
			Scope: "eosio",
			Code:  "eosio",
			Table: "bps",
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
		var producers producers
		err := json.Unmarshal(response.Rows, &producers)
		errorCheck("json marshaling", err)

		fmt.Println("List of producers registered to receive votes:")
		for _, p := range producers {
			fmt.Printf("- %s (key: %s)\n", p["name"], p["block_signing_key"])
		}
		fmt.Printf("Total of %d registered producers\n", len(producers))

	}
}

func init() {
	systemCmd.AddCommand(systemListProducersCmd)

	systemListProducersCmd.Flags().BoolP("json", "j", false, "return producers info in json")

	for _, flag := range []string{"json"} {
		if err := viper.BindPFlag("vote-list-cmd-"+flag, systemListProducersCmd.Flags().Lookup(flag)); err != nil {
			panic(err)
		}
	}
}
