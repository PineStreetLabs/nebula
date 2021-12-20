package account

import (
	"github.com/PineStreetLabs/nebula/networks"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// An Account is defined by a public and private key pair.
// An address is associated with an account based on context (e.g. users, validators).
// The cosmos-sdk `BaseAccount` type is used to encapsulate the account model.
// An account might have extended functionality based on the app chain.

type Account struct {
	address   *Address
	publicKey cryptotypes.PubKey
	sequence  uint64
	number    uint64
}

func (a Account) GetAddress() sdk.Address {
	return a.address
}

func (a Account) GetPubKey() cryptotypes.PubKey {
	return a.publicKey
}

func (a Account) GetAccountNumber() uint64 {
	return a.number
}

func (a Account) GetSequence() uint64 {
	return a.sequence
}

// FromPublicKey creates an account address using app configuration and a public key.
func FromPublicKey(cfg *networks.Params, pk cryptotypes.PubKey) (*Account, error) {
	buf := pk.Address().Bytes()
	hrp := cfg.AccountHRP()

	if err := cfg.VerifyAddressFormat(buf); err != nil {
		return nil, err
	}

	return &Account{
		address: &Address{
			data: buf,
			hrp:  hrp,
		},
		publicKey: pk,
		sequence:  0,
		number:    0,
	}, nil
}
