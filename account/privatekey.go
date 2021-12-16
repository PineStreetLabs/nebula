package account

import "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

// NewPrivateKey uses the underlying cosmos-sdk/crypto to generate a private key.
func NewPrivateKey() *secp256k1.PrivKey {
	return secp256k1.GenPrivKey()
}
