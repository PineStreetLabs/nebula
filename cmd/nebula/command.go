package main

import "github.com/urfave/cli"

var newAccountCommand = cli.Command{
	Name:     "new",
	Category: "wallet",
	Usage:    "Create a new account.",
	Description: `
	Creates a new account depending on the network parameters.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "network",
			Usage:    "network parameters",
			Required: true,
		},
		cli.StringFlag{
			Name:     "from",
			Usage:    "base64 encoded 32 byte seed",
			Required: false,
		},
	},
	Action: newAccount,
}
