package account

import (
	"encoding/base64"
	"errors"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

var errInvalidPkLen = errors.New("invalid public key length")

// ParseSecp256k1PublicKey accepts a hex-encoded string and returns a secp256k1 public key.
func ParseSecp256k1PublicKey(pk string) (*secp256k1.PubKey, error) {
	key, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		return nil, err
	}

	if len(key) != secp256k1.PubKeySize {
		return nil, errInvalidPkLen
	}

	// TODO verify
	return &secp256k1.PubKey{
		Key: key,
	}, nil
}

func ParseEd25519PublicKey(pk string) (*ed25519.PubKey, error) {
	key, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		return nil, err
	}

	if len(key) != ed25519.PubKeySize {
		return nil, errInvalidPkLen
	}

	// TODO verify
	return &ed25519.PubKey{
		Key: key,
	}, nil
}
