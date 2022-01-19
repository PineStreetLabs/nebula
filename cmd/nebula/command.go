package main

import "github.com/urfave/cli"
import "github.com/PineStreetLabs/nebula/cmd/nebula/common"

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
	Flags: append([]cli.Flag{
		cli.StringFlag{
			Name:     "recipient",
			Usage:    "recipient's address",
			Required: true,
		},
		cli.Uint64Flag{
			Name:     "amount",
			Usage:    "unsigned integer amount to send to the recipient",
			Required: true,
		},
	}, common.TxFlags...),
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
