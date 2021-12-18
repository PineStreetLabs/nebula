package account

import (
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

// An Account is defined by a public and private key pair.
// An address is associated with an account based on context (e.g. users, validators).
// The cosmos-sdk `BaseAccount` type is used to encapsulate the account model.

// FromPublicKey creates an account address using app configuration and a public key.
func FromPublicKey(cfg *sdk.Config, pk types.PubKey) (sdk.AccAddress, error) {
	buf := pk.Address().Bytes()
	hrp := cfg.GetBech32AccountAddrPrefix()
	addr, err := bech32.ConvertAndEncode(hrp, buf)
	if err != nil {
		return nil, err
	}
	// todo, network specific verifyAddressFormat...
	//err = VerifyAddressFormat(bz)

	return sdk.AccAddress(addr), nil
}
