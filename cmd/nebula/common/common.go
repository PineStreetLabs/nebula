package common

import (
	"errors"
	"fmt"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/urfave/cli"
)

var errUnsupportedNetwork = errors.New("unsupported network")

func GetNetworkConfig(ctx *cli.Context) (*networks.Params, error) {
	switch network := ctx.GlobalString("network"); network {
	case networks.Cosmos:
		return networks.GetCosmosCfg(), nil
	case networks.Umee:
		return networks.GetUmeeCfg(), nil
	case "":
		return nil, errors.New("missing network")
	default:
		return nil, fmt.Errorf("%w : %s", errUnsupportedNetwork, network)
	}
}

var TxFlags = []cli.Flag{
	// account info
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
	// network info
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
		Name:     "memo",
		Usage:    "a note or comment to send with the transaction",
		Required: true,
	},
}
