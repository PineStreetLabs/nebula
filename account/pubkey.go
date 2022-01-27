package account

import (
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// ParsePublicKey accepts a hex-encoded string and returns a secp256k1 public key.
func ParsePublicKey(pk string) (*secp256k1.PubKey, error) {
	key, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		return nil, err
	}
	// verify..
	return &secp256k1.PubKey{
		Key: key,
	}, nil
}
