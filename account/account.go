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
	// Each account is identified using an Address derived from a public key.
	// For now, we assume the Account is always represented by a user address (AccAddress).
	address   *Address
	publicKey cryptotypes.PubKey
	// sequence is the Account Sequence. The sequence is an incremental nonce used for replay-protection per network.
	sequence uint64
	// number is the Account Number. This is a globally defined nonce that is associated with an account.
	// The Account Number is initialized per network when it is "initialized" on the network (e.g. first receive).
	number uint64
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

// SetSequence assigns a new sequence to the Account.
func (a *Account) SetSequence(accSeq uint64) {
	a.sequence = accSeq
}

// FromPublicKey creates an account address using app configuration and a public key.
func FromPublicKey(cfg *networks.Params, pk cryptotypes.PubKey, accNum, accSeq uint64) (*Account, error) {
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
		sequence:  accSeq,
		number:    accNum,
	}, nil
}
