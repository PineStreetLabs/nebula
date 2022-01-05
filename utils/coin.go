package utils

import (
	"github.com/PineStreetLabs/nebula/networks"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewCoin returns a Coin with the provided network's denom.
func NewCoin(cfg *networks.Params, value sdk.Int) sdk.Coin {
	return sdk.Coin{
		Denom:  cfg.Denom(),
		Amount: value,
	}
}
