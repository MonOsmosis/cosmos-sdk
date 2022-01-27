package main

import (
	"os"

	"github.com/MonOsmosis/cosmos-sdk/server"
	svrcmd "github.com/MonOsmosis/cosmos-sdk/server/cmd"
	"github.com/MonOsmosis/cosmos-sdk/simapp"
	"github.com/MonOsmosis/cosmos-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
