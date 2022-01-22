package umee

import "github.com/urfave/cli"
import "github.com/PineStreetLabs/nebula/cmd/nebula/common"

var LendAssetCommand = cli.Command{
	Name:     "lend_asset",
	Category: "umee",
	Usage:    "Create a lend asset transaction.",
	Description: `
	An action performed by a lender.
	Creates a single transaction and returns the serialized transaction in bytes.
	`,
	Flags: append([]cli.Flag{
		cli.Uint64Flag{
			Name:     "amount",
			Usage:    "unsigned integer amount to lend",
			Required: true,
		},
	}, common.TxFlags...),
	Action: lendAsset,
}

var WithdrawAssetCommand = cli.Command{
	Name:     "withdraw_asset",
	Category: "umee",
	Usage:    "Create a withdraw asset transaction.",
	Description: `
    A lender's request to withdraw a previously lent asset
	Creates a single transaction and returns the serialized transaction in bytes.
	`,
	Flags: append([]cli.Flag{
		cli.Uint64Flag{
			Name:     "amount",
			Usage:    "unsigned integer amount to withdraw by the lender",
			Required: true,
		},
	}, common.TxFlags...),
	Action: withdrawAsset,
}

var SetCollateralCommand = cli.Command{
	Name:     "set_collateral",
	Category: "umee",
	Usage:    "Create a set collateral transaction.",
	Description: `
	A lender's request to enable or disable a token type in their possession as collateral
	Creates a single transaction and returns the serialized transaction in bytes.
	`,
	Flags: append([]cli.Flag{
		cli.BoolFlag{
			Name:     "enabled",
			Usage:    "true to enable tokens as collateral",
			Required: true,
		},
	}, common.TxFlags...),
	Action: setCollateral,
}

var RepayAssetCommand = cli.Command{
	Name:     "repay_asset",
	Category: "umee",
	Usage:    "Create a repay asset transaction.",
	Description: `
	An action performed by a borrower.
	Creates a single transaction and returns the serialized transaction in bytes.
	`,
	Flags: append([]cli.Flag{
		cli.Uint64Flag{
			Name:     "amount",
			Usage:    "the unsigned integer amount to repay",
			Required: true,
		},
	}, common.TxFlags...),
	Action: repayAsset,
}
