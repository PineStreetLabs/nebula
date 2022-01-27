package keychain

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/go-bip39"
)

// Keychain.go implements the BIP32 HD standard.
// We rely on the cosmos-sdk crypto/hd's hdkeychain as a dependency.
// This package solely implements the plumbing around bip32 missing in the ecosystem,
// including the derivation path specs.

var errSeedLen = errors.New("seed is not 32 bytes")

// ExtendedKey contains the necessary key information for BIP32 key derivation.
type ExtendedKey struct {
	key       [32]byte
	chaincode [32]byte
}

// FromSeed computes a new master key for BIP32 HD key derivation.
func FromSeed(seed []byte) (*ExtendedKey, error) {
	// https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#master-key-generation
	// Generate a seed byte sequence S of a chosen length (between 128 and 512 bits; 256 bits is advised)
	if seedLen := len(seed); seedLen < 16 || seedLen > 64 {
		return nil, errSeedLen
	}

	masterSecret, chaincode := hd.ComputeMastersFromSeed(seed)

	return &ExtendedKey{
		key:       masterSecret,
		chaincode: chaincode,
	}, nil
}

// Derive does child key derivation with the provided extended key and BIP44 path.
// The function returns the child in bytes.
func Derive(extKey *ExtendedKey, path *hd.BIP44Params) ([]byte, error) {
	// The underlying function requires a strict BIP44 path.
	return hd.DerivePrivateKeyForPath(extKey.key, extKey.chaincode, path.String())
}

// FromMnemonic creates a new extended key from a BIP39 mnemonic phrase.
func FromMnemonic(mnemonic, bip39Passphrase string) (*ExtendedKey, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, bip39Passphrase)
	if err != nil {
		return nil, err
	}

	return FromSeed(seed)
}
