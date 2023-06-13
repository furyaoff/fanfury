package main

import (
	"github.com/furyaoff/fanfury/app/params"
	"os"

	"github.com/furyaoff/fanfury/app"
	"github.com/furyaoff/fanfury/cmd/fanfuryd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	params.SetAddressPrefixes()

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
