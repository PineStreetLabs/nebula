package main

import "github.com/urfave/cli"

var newAccountCommand = cli.Command{
	Name:     "account",
	Category: "wallet",
	Usage:    "Create a new account.",
	Description: `
	Creates a new account depending on the network parameters.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "from_sk",
			Usage:    "base64 encoded secret key",
			Required: false,
		},
		cli.StringFlag{
			Name:     "from_mnemonic",
			Usage:    "space separated string of BIP39 mnemonics",
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
			Name:     "recipient",
			Usage:    "recipient's address",
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
		cli.Uint64Flag{
			Name:     "acc_number",
			Usage:    "account number",
			Required: true,
		},
		cli.Uint64Flag{
			Name:     "acc_sequence",
			Usage:    "account sequence",
			Required: true,
		},
		cli.StringFlag{
			Name:     "memo",
			Usage:    "a note or comment to send with the transaction",
			Required: true,
		},
		cli.Uint64Flag{
			Name:     "amount",
			Usage:    "unsigned integer amount to send to the recipient",
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
			Name:     "tx_hex",
			Usage:    "Serialized transaction in hex",
			Required: true,
		},
	},
	Action: broadcastTx,
}

var balanceCommand = cli.Command{
	Name:     "balance",
	Category: "data",
	Usage:    "<address>",
	Description: `
	Checks balance for an account.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "address",
			Required: true,
		},
	},
	Action: queryBalance,
}

var accountCommand = cli.Command{
	Name:     "account_info",
	Category: "data",
	Usage:    "<address>",
	Description: `
	Checks details for an account.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "address",
			Required: true,
		},
	},
	Action: queryAccount,
}

var bestBlockHeightCommand = cli.Command{
	Name:     "bestblockheight",
	Category: "data",
	Description: `
	Gets latest blockheight for a network.
	`,
	Action: queryBestBlockheight,
}

var blockByHeightCommand = cli.Command{
	Name:     "blockbyheight",
	Category: "data",
	Usage:    "<height>",
	Description: `
	Gets a block by height.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "height",
			Required: true,
		},
	},
	Action: queryBlockByHeight,
}

var blockByHashCommand = cli.Command{
	Name:     "blockbyhash",
	Category: "data",
	Usage:    "<hash>",
	Description: `
	Gets a block by hash.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "hash",
			Required: true,
		},
	},
	Action: queryBlockByHash,
}
