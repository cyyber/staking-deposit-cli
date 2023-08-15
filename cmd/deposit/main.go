package main

import (
	"os"

	"github.com/cyyber/staking-deposit-cli/cmd/deposit/existingmnemonic"
	"github.com/cyyber/staking-deposit-cli/cmd/deposit/newmnemonic"
	"github.com/urfave/cli/v2"
)

var depositCommands []*cli.Command

func main() {
	app := &cli.App{
		Commands: depositCommands,
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func init() {
	depositCommands = append(depositCommands, existingmnemonic.Commands...)
	depositCommands = append(depositCommands, newmnemonic.Commands...)
}
