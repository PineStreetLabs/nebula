package account

import (
	"bytes"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"gopkg.in/yaml.v2"
)

// Address implements the sdk.Address type.
type Address struct {
	data []byte
	hrp  string
}

// AddressFromString returns an Address from its string encoding.
func AddressFromString(addr string) (*Address, error) {
	hrp, data, err := bech32.DecodeAndConvert(addr)
	if err != nil {
		return nil, err
	}

	return &Address{
		data: data,
		hrp:  hrp,
	}, nil
}

// Equals returns boolean for whether two Addresses are Equal
func (a Address) Equals(a2 sdk.Address) bool {
	if a.Empty() && a2.Empty() {
		return true
	}

	return bytes.Equal(a.Bytes(), a2.Bytes())
}

// Empty returns boolean for whether an Address is empty
func (a Address) Empty() bool {
	return a.data == nil || len(a.data) == 0
}

// Marshal returns the raw address bytes. It is needed for protobuf
// compatibility.
func (a Address) Marshal() ([]byte, error) {
	return a.data, nil
}

// Unmarshal sets the address to the given data. It is needed for protobuf
// compatibility.
func (a *Address) Unmarshal(data []byte) error {
	a.data = data
	return nil
}

// MarshalJSON marshals to JSON using Bech32.
func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// MarshalYAML marshals to YAML using Bech32.
func (a Address) MarshalYAML() (interface{}, error) {
	return a.String(), nil
}

// UnmarshalJSON unmarshals from JSON assuming Bech32 encoding.
func (a *Address) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)

	if err != nil {
		return err
	}
	if s == "" {
		*a = Address{}
		return nil
	}

	a2, err := sdk.GetFromBech32(s, a.hrp)
	if err != nil {
		return err
	}

	a.data = a2
	return nil
}

// UnmarshalYAML unmarshals from JSON assuming Bech32 encoding.
func (a *Address) UnmarshalYAML(data []byte) error {
	var s string
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*a = Address{}
		return nil
	}

	hrp, data, err := bech32.DecodeAndConvert(s)
	if err != nil {
		return err
	}

	*a = Address{
		data: data,
		hrp:  hrp,
	}
	return nil
}

// Bytes returns the raw address bytes.
func (a Address) Bytes() []byte {
	return a.data
}

// String implements the Stringer interface.
func (a Address) String() string {
	if a.Empty() {
		return ""
	}

	addr, _ := bech32.ConvertAndEncode(a.hrp, a.data)
	return addr
}

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (a Address) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(a.String()))
	case 'p':
		s.Write([]byte(fmt.Sprintf("%p", a)))
	default:
		s.Write([]byte(fmt.Sprintf("%X", a.data)))
	}
}
