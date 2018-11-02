package cmd

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	eos "github.com/eosforce/goeosforce"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var toolsNamesCmd = &cobra.Command{
	Use:   "names [value]",
	Short: "Convert a value to and from name-encoded strings",
	Long: `EOS name encoding creates strings or up to 12 characters out of uint64 values.

This command auto-detects encoding and converts it to different encodings.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		showFrom := map[string]uint64{}

		baseHex, err := hex.DecodeString(input)
		if err == nil {
			if len(baseHex) == 8 {
				showFrom["hex"] = binary.LittleEndian.Uint64(baseHex)
				showFrom["hex_be"] = binary.BigEndian.Uint64(baseHex)
			} else if len(baseHex) == 4 {
				showFrom["hex"] = uint64(binary.LittleEndian.Uint32(baseHex))
				showFrom["hex_be"] = uint64(binary.BigEndian.Uint32(baseHex))
			}
		}

		fromName, err := eos.StringToName(input)
		if err == nil {
			showFrom["name"] = fromName
		}

		fromUint64, err := strconv.ParseUint(input, 10, 64)
		if err == nil {
			showFrom["uint64"] = fromUint64
		}

		someFound := false
		rows := []string{"| from \\ to | hex | hex_be | name | uint64", "| --------- | --- | ------ | ---- | ------ |"}
		for _, from := range []string{"hex", "hex_be", "name", "uint64"} {
			val, found := showFrom[from]
			if !found {
				continue
			}
			someFound = true

			row := []string{from}
			for _, to := range []string{"hex", "hex_be", "name", "uint64"} {

				cnt := make([]byte, 8)
				switch to {
				case "hex":
					binary.LittleEndian.PutUint64(cnt, val)
					row = append(row, hex.EncodeToString(cnt))
				case "hex_be":
					binary.BigEndian.PutUint64(cnt, val)
					row = append(row, hex.EncodeToString(cnt))

				case "name":
					row = append(row, eos.NameToString(val))

				case "uint64":
					row = append(row, strconv.FormatUint(val, 10))
				}
			}
			rows = append(rows, "| "+strings.Join(row, " | ")+" |")
		}

		if !someFound {
			fmt.Printf("Couldn't decode %q with any of these methods: hex, hex_be, name, uint64\n", input)
			os.Exit(1)
		}

		fmt.Println("")
		fmt.Println(columnize.SimpleFormat(rows))
		fmt.Println("")
	},
}

func init() {
	toolsCmd.AddCommand(toolsNamesCmd)
}
