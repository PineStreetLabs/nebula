package account

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

var errInvalidPkLen = errors.New("invalid public key length")

// Secp256k1PublicKey accepts a byte slice and returns a secp256k1 public key.
func Secp256k1PublicKey(pk []byte) (*secp256k1.PubKey, error) {
	if len(pk) != secp256k1.PubKeySize {
		return nil, errInvalidPkLen
	}

	// TODO verify
	return &secp256k1.PubKey{
		Key: pk,
	}, nil
}

// Ed25519PublicKey accepts a byte slice and returns a ed25519 public key.
func Ed25519PublicKey(pk []byte) (*ed25519.PubKey, error) {
	if len(pk) != ed25519.PubKeySize {
		return nil, errInvalidPkLen
	}

	// TODO verify
	return &ed25519.PubKey{
		Key: pk,
	}, nil
}
