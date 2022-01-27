package umee

import (
	"github.com/PineStreetLabs/nebula/cmd/nebula/common"
	"github.com/urfave/cli"
)

// LendAssetCommand exposes functionality to lend assets on Umee.
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

// WithdrawAssetCommand exposes functionality to withdraw assets on Umee.
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

// SetCollateralCommand exposes functionality to set collateral on Umee.
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

// RepayAssetCommand exposes functionality to repay assets on Umee.
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
