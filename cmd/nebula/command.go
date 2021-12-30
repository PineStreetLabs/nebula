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

var newBankSendCommand = cli.Command{
	Name:     "bank_send",
	Category: "wallet",
	Usage:    "Create a new bank send transaction.",
	Description: `
	Creates a single bank send transaction and returns the serialized transaction in bytes.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "network",
			Usage:    "network parameters",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "fee",
			Usage:    "the maximum amount the user is willing to pay in fees",
			Required: true,
		},
		cli.Uint64Flag{
			Name:     "gas_limit",
			Usage:    "option chosen by the users for how to calculate how much gas they will need to pay",
			Required: true,
		},
		cli.Uint64Flag{
			Name:     "timeout_height",
			Usage:    "block height until which the transaction is valid",
			Required: true,
		},
		cli.StringFlag{
			Name:     "private_key",
			Usage:    "private key to sign transaction",
			Required: true,
		},
		cli.StringFlag{
			Name:     "memo",
			Usage:    "a note or comment to send with the transaction",
			Required: true,
		},
	},
	Action: newBankSend,
}

var broadcastTxCommand = cli.Command{
	Name:     "broadcast_tx",
	Category: "wallet",
	Usage:    "Broadcast a transaction",
	Description: `
	Broadcasts a transaction.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "network",
			Usage:    "network parameters",
			Required: true,
		},
		cli.StringFlag{
			Name:     "tx_hex",
			Usage:    "Serialized transaction in hex",
			Required: true,
		},
		cli.StringFlag{
			Name:     "host",
			Usage:    "host name",
			Required: true,
		},
		cli.IntFlag{
			Name:     "port",
			Usage:    "port number",
			Required: true,
		},
	},
	Action: broadcastTx,
}
