package utils

import (
	"math/big"

	"github.com/PineStreetLabs/nebula/networks"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewCoin returns a Coin with the provided network's denom.
func NewCoin(cfg *networks.Params, value *big.Int) sdk.Coin {
	return sdk.Coin{
		Denom:  cfg.Denom(),
		Amount: sdk.NewIntFromBigInt(value),
	}
}
