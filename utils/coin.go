package utils

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewCoinFromBigInt returns a Coin from big.Int.
func NewCoinFromBigInt(denom string, value *big.Int) sdk.Coin {
	return sdk.Coin{
		Denom:  denom,
		Amount: sdk.NewIntFromBigInt(value),
	}
}

// NewCoinFromUint64 returns a Coin from uint64.
func NewCoinFromUint64(denom string, value uint64) sdk.Coin {
	return sdk.Coin{
		Denom:  denom,
		Amount: sdk.NewIntFromUint64(value),
	}
}
