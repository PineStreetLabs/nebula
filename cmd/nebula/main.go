package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "nebula"
	app.Usage = "Gateway to the Cosmos."

	app.Commands = []cli.Command{
		broadcastTxCommand,
		newAccountCommand,
		newBankSendCommand,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
