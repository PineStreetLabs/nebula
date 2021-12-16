package account

import (
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

// FromPublicKey creates an account address using app configuration and a public key.
func FromPublicKey(cfg *sdk.Config, pk types.PubKey) (sdk.Address, error) {
	buf := pk.Address().Bytes()
	hrp := cfg.GetBech32AccountAddrPrefix()
	addr, err := bech32.ConvertAndEncode(hrp, buf)
	if err != nil {
		return nil, err
	}
	return sdk.AccAddressFromBech32(addr)
}
