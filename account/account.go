package account

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
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

// JSON is a helper struct for serialization.
type JSON struct {
	Address   string `json:"address"`
	PublicKey string `json:"publickey"`
	Sequence  uint64 `json:"sequence"`
	Number    uint64 `json:"number"`
}

// MarshalJSON implements the marshaller interface.
func (a *Account) MarshalJSON() ([]byte, error) {
	var pk string

	if a.publicKey == nil {
		pk = ""
	} else {
		pk = fmt.Sprintf("%x", a.publicKey.Bytes())
	}
	return json.Marshal(JSON{
		Address:   a.address.String(),
		PublicKey: pk,
		Sequence:  a.sequence,
		Number:    a.number,
	})
}

// UnmarshalJSON implements the marshaller interface.
func (a *Account) UnmarshalJSON(b []byte) error {
	temp := &JSON{}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	address, err := AddressFromString(temp.Address)
	if err != nil {
		return err
	}

	pkBytes, err := hex.DecodeString(temp.PublicKey)
	if err != nil {
		return err
	}

	pk := &secp256k1.PubKey{
		Key: pkBytes,
	}

	*a = Account{}

	*a = Account{
		publicKey: pk,
		number:    temp.Number,
		sequence:  temp.Sequence,
		address:   address,
	}

	return nil
}

// GetAddress returns Account's address.
func (a Account) GetAddress() sdk.Address {
	return a.address
}

// GetPubKey returns Account's public key.
func (a Account) GetPubKey() cryptotypes.PubKey {
	return a.publicKey
}

// GetAccountNumber returns Account's account nonce (number).
func (a Account) GetAccountNumber() uint64 {
	return a.number
}

// GetSequence returns Account's sequence.
func (a Account) GetSequence() uint64 {
	return a.sequence
}

// SetSequence assigns a new sequence to the Account.
func (a *Account) SetSequence(accSeq uint64) {
	a.sequence = accSeq
}

var errMissingPublicKey = errors.New("missing public key")

// NewAccount creates a new account.
func NewAccount(address string, publickey cryptotypes.PubKey, accNum, accSeq uint64) (*Account, error) {
	addr, err := AddressFromString(address)
	if err != nil {
		return nil, err
	}

	if publickey == nil || len(publickey.Bytes()) == 0 {
		return nil, errMissingPublicKey
	}
	// Make sure publickey is P256 curve..

	return &Account{
		address:   addr,
		publicKey: publickey,
		number:    accNum,
		sequence:  accSeq,
	}, nil
}

// ValidatorFromPublicKey creates an account with a validator address.
func NewValidatorAccount(cfg *networks.Params, pk cryptotypes.PubKey, accNum, accSeq uint64) (*Account, error) {
	hrp := cfg.ValidatorHRP()
	return fromPublicKey(hrp, cfg, pk, accNum, accSeq)
}

// FromPublicKey creates an account address using app configuration and a public key.
func NewUserAccount(cfg *networks.Params, pk cryptotypes.PubKey, accNum, accSeq uint64) (*Account, error) {
	hrp := cfg.AccountHRP()
	return fromPublicKey(hrp, cfg, pk, accNum, accSeq)
}

func NewConsensusAccount(cfg *networks.Params, pk cryptotypes.PubKey, accNum, accSeq uint64) (*Account, error) {
	hrp := cfg.ConsensusHRP()
	return fromPublicKey(hrp, cfg, pk, accNum, accSeq)
}

func fromPublicKey(hrp string, cfg *networks.Params, pk cryptotypes.PubKey, accNum, accSeq uint64) (*Account, error) {
	if pk == nil || len(pk.Bytes()) == 0 {
		return nil, errMissingPublicKey
	}

	buf := pk.Address().Bytes()

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

// ToValidatorAddress accepts an Account and attempts to convert the address into an sdk.ValAddress type.
func ToValidatorAddress(cfg *networks.Params, account *Account) (sdk.ValAddress, error) {
	hrp := cfg.ValidatorHRP()

	if hrp != account.address.hrp {
		return nil, errors.New("account is not a validator")
	}

	bz := account.address.data

	if err := cfg.VerifyAddressFormat(bz); err != nil {
		return nil, err
	}

	return sdk.ValAddress(bz), nil
}
