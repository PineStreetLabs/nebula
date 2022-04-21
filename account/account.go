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
	"github.com/gogo/protobuf/proto"
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
	number     uint64
	privateKey cryptotypes.PrivKey
}

// JSON is a helper struct for serialization.
type JSON struct {
	Address    string `json:"address"`
	PublicKey  string `json:"publickey"`
	Sequence   uint64 `json:"sequence"`
	Number     uint64 `json:"number"`
	PrivateKey string `json:"privatekey,omitempty"`
}

// MarshalJSON implements the marshaller interface.
func (a *Account) MarshalJSON() ([]byte, error) {
	var pk, sk string

	if a.publicKey != nil {
		pk = fmt.Sprintf("%x", a.publicKey.Bytes())
	}

	if a.privateKey != nil {
		sk = fmt.Sprintf("%x", a.privateKey.Bytes())
	}

	return json.Marshal(JSON{
		Address:    a.address.String(),
		PublicKey:  pk,
		Sequence:   a.sequence,
		Number:     a.number,
		PrivateKey: sk,
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

	var pk cryptotypes.PubKey
	if len(temp.PublicKey) != 0 {
		pkBytes, err := hex.DecodeString(temp.PublicKey)
		if err != nil {
			return err
		}

		pk = &secp256k1.PubKey{
			Key: pkBytes,
		}

	} else {
		pk = &secp256k1.PubKey{}
	}

	var sk cryptotypes.PrivKey
	if len(temp.PrivateKey) != 0 {
		skBytes, err := hex.DecodeString(temp.PrivateKey)
		if err != nil {
			return err
		}

		sk = &secp256k1.PrivKey{
			Key: skBytes,
		}
	} else {
		sk = &secp256k1.PrivKey{}
	}

	*a = Account{
		number:     temp.Number,
		sequence:   temp.Sequence,
		address:    address,
		privateKey: sk,
		publicKey:  pk,
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

var errInvalidKey = errors.New("invalid key")
var errMissingPublicKey = errors.New("missing public key")

// NewAccount creates a new account.
func NewAccount(address string, publickey cryptotypes.PubKey, accNum, accSeq uint64) (*Account, error) {
	addr, err := AddressFromString(address)
	if err != nil {
		return nil, err
	}

	return &Account{
		address:   addr,
		publicKey: publickey,
		number:    accNum,
		sequence:  accSeq,
	}, nil
}

// Key is an interface for describing either a private or public key.
type Key interface {
	proto.Message
	Bytes() []byte
}

// NewValidatorAccount creates an Account for the validator context.
func NewValidatorAccount(cfg *networks.Params, key Key, accNum, accSeq uint64) (*Account, error) {
	hrp := cfg.ValidatorHRP()
	return newAccount(hrp, cfg, key, accNum, accSeq)
}

// NewUserAccount creates an Account for the user/account context.
func NewUserAccount(cfg *networks.Params, key Key, accNum, accSeq uint64) (acc *Account, err error) {
	hrp := cfg.AccountHRP()
	return newAccount(hrp, cfg, key, accNum, accSeq)
}

// NewConsensusAccount creates an Account for the consensus context.
func NewConsensusAccount(cfg *networks.Params, key Key, accNum, accSeq uint64) (*Account, error) {
	hrp := cfg.ConsensusHRP()
	return newAccount(hrp, cfg, key, accNum, accSeq)
}

// newAccount creates a new account from either a public or private key.
// If it is a private key, we include the key material in the returned Account.
func newAccount(hrp string, cfg *networks.Params, key Key, accNum, accSeq uint64) (acc *Account, err error) {
	switch k := key.(type) {
	case cryptotypes.PrivKey:
		acc, err = fromPublicKey(hrp, cfg, k.PubKey(), accNum, accSeq)
		if err != nil {
			return nil, err
		}

		acc.privateKey = k
	case cryptotypes.PubKey:
		acc, err = fromPublicKey(hrp, cfg, k, accNum, accSeq)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errInvalidKey
	}

	return acc, nil
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
