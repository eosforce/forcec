package main

import (
	// Load all contracts here, so we can always read and decode
	// transactions with those contracts.
	_ "github.com/eosforce/goeosforce/msig"
	_ "github.com/eosforce/goeosforce/system"
	_ "github.com/eosforce/goeosforce/token"

	"github.com/eosforce/forcec/eosc/cmd"
)

var version = "dev"

func init() {
	cmd.Version = version
}

func main() {
	cmd.Execute()
}
