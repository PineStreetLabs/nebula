package account

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// NewPrivateKey uses the underlying cosmos-sdk/crypto to generate a private key.
func NewPrivateKey() *secp256k1.PrivKey {
	return secp256k1.GenPrivKey()
}

// PrivateKeyFromHex accepts a hex-encoded string and returns the sepc256k1 private key.
func PrivateKeyFromHex(sk string) (*secp256k1.PrivKey, error) {
	buf, err := hex.DecodeString(sk)
	if err != nil {
		return nil, err
	}

	return &secp256k1.PrivKey{
		Key: buf,
	}, nil
}
