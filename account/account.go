package account

import (
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

// An Account is defined by a public and private key pair.
// An address is associated with an account based on context (e.g. users, validators).
// The cosmos-sdk `BaseAccount` type is used to encapsulate the account model.
// An account might have extended functionality based on the app chain.

type Account struct {
	address   sdk.AccAddress
	publicKey cryptotypes.PubKey
	sequence  uint64
	number    uint64
}

func (a Account) GetAddress() sdk.AccAddress {
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

// todo, various types of accounts..
// todo validation
// todo

// FromPublicKey creates an account address using app configuration and a public key.
func FromPublicKey(cfg *networks.Params, pk cryptotypes.PubKey) (*Account, error) {
	buf := pk.Address().Bytes()
	hrp := cfg.AccountHRP()
	addr, err := bech32.ConvertAndEncode(hrp, buf)
	if err != nil {
		return nil, err
	}
	// todo, network specific verifyAddressFormat...
	//err = VerifyAddressFormat(bz)

	return &Account{
		address:   sdk.AccAddress(addr),
		publicKey: pk,
		sequence:  0,
		number:    0,
	}, nil
}

func FromAccount(acc client.Account) *Account {
	return &Account{
		address:   acc.GetAddress(),
		publicKey: acc.GetPubKey(),
		sequence:  acc.GetSequence(),
		number:    acc.GetAccountNumber(),
	}
}
