package main

import (
	"github.com/PineStreetLabs/nebula/cmd/nebula/common"
	"github.com/urfave/cli"
)

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
		cli.BoolFlag{
			Name:     "ledger",
			Usage:    "use Ledger device",
			Required: false,
		},
		cli.IntFlag{
			Name:     "ledger_account",
			Usage:    "account to use on Ledger",
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

var newMultiSigAccountCommand = cli.Command{
	Name:     "multisig_account",
	Category: "wallet",
	Usage:    "Create a new multisig account.",
	Description: `
	Creates a new multisig account.
	`,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:     "threshold",
			Usage:    "m-of-n threshold",
			Required: true,
		},
		cli.StringSliceFlag{
			Name:     "publickey",
			Usage:    "list of public keys",
			Required: true,
		},
	},
	Action: newMultiSig,
}

var newBankSendCommand = cli.Command{
	Name:     "bank_send",
	Category: "wallet",
	Usage:    "Create a Bank module MsgSend message.",
	Description: `
	Create a Bank module MsgSend message.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "sender",
			Usage:    "senders's address",
			Required: true,
		},
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
	},
	Action: newBankSend,
}

var signTxCommand = cli.Command{
	Name:        "sign_tx",
	Category:    "wallet",
	Usage:       "Sign a serialized transaction.",
	Description: "Sign a serialized transaction.",
	Flags: append([]cli.Flag{
		cli.StringFlag{
			Name:     "tx",
			Usage:    "hex-encoded transaction",
			Required: true,
		},
		cli.StringFlag{
			Name:     "private_key",
			Usage:    "private key to sign transaction",
			Required: true,
		},
	}, common.SignTxFlags...),
	Action: signTx,
}

var partialSignTxCommand = cli.Command{
	Name:        "partial_sign_tx",
	Category:    "wallet",
	Usage:       "Sign a serialized transaction using a Ledger device.",
	Description: "Sign a serialized transaction using a Ledger device.",
	Flags: append([]cli.Flag{
		cli.StringFlag{
			Name:     "tx",
			Usage:    "hex-encoded transaction",
			Required: true,
		},
		cli.IntFlag{
			Name:     "ledger_account",
			Usage:    "account to use on Ledger",
			Required: true,
		},
	}, common.SignTxFlags...),
	Action: partialSignTx,
}

var combineTxCommand = cli.Command{
	Name:        "combine_signatures",
	Category:    "wallet",
	Usage:       "Combines signatures for a multisignature account and finalizes a transaction.",
	Description: "Combines signatures for a multisignature account and finalizes a transaction.",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "tx",
			Usage:    "hex-encoded transaction",
			Required: true,
		},
		cli.StringSliceFlag{
			Name:     "signature",
			Usage:    "json signatures",
			Required: true,
		},
		cli.StringFlag{
			Name:     "tx",
			Usage:    "hex-encoded transaction",
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
			Name:     "multisig_account",
			Usage:    "JSON multisig account",
			Required: true,
		},
	},
	Action: partialSignTx,
}

var newTxCommand = cli.Command{
	Name:        "new_tx",
	Category:    "wallet",
	Usage:       "Combines a slice of messages into a new transaction.",
	Description: "Combines a slice of messages into a new transaction.",
	Flags: append([]cli.Flag{
		cli.StringSliceFlag{
			Name:     "messages",
			Required: true,
		},
	}, common.NewTxFlags...),
	Action: newTx,
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

var transactionCommand = cli.Command{
	Name:     "transaction",
	Category: "data",
	Usage:    "<txid>",
	Description: `
	Gets a transaction by its ID.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "txid",
			Required: true,
		},
	},
	Action: queryTransaction,
}
