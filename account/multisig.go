package account

import (
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/types"
)

// NewMultiSigAccount creates a "naive-multisig" public key that composes sub-publickey objects.
func NewMultiSigAccount(cfg *networks.Params, threshold int, pubkeys []types.PubKey, accNum, accSeq uint64) (*Account, error) {
	key := multisig.NewLegacyAminoPubKey(threshold, pubkeys)

	return &Account{
		address: &Address{
			data: key.Address().Bytes(),
			hrp:  cfg.AccountHRP(),
		},
		publicKey: key,
		number:    accNum,
		sequence:  accSeq,
	}, nil
}
