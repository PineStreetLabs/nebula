package account

import (
	"encoding/base64"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

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
