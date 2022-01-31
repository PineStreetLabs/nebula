package utils

import (
	"math/big"

	"github.com/PineStreetLabs/nebula/networks"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewCoinFromBigInt returns a Coin from big.Int.
func NewCoinFromBigInt(cfg *networks.Params, value *big.Int) sdk.Coin {
	return sdk.Coin{
		Denom:  cfg.Denom(),
		Amount: sdk.NewIntFromBigInt(value),
	}
}

// NewCoinFromUint64 returns a Coin from uint64.
func NewCoinFromUint64(cfg *networks.Params, value uint64) sdk.Coin {
	return sdk.Coin{
		Denom:  cfg.Denom(),
		Amount: sdk.NewIntFromUint64(value),
	}
}
