package networks

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// Params includes the scope and configuration of each network.
// In the cosmos-sdk, this is akin to the sdk.Config type.
// We choose not to use sdk.Config because of its use of init and how it is used globally in the cosmos-sdk library.
// Params tightly scopes relevant app chain configurations to work neatly across this library.
type Params struct {
	denom               string
	accountHRP          string
	validatorHRP        string
	consensusHRP        string
	VerifyAddressFormat func(b []byte) error
	encodingConfig      EncodingConfig
}

// Denom gets the denom.
func (p Params) Denom() string {
	return p.denom
}

// AccountHRP gets the human readable prefix for an Account's address.
func (p Params) AccountHRP() string {
	return p.accountHRP
}

// ValidatorHRP gets the human readable prefix for an Account's validator address.
func (p Params) ValidatorHRP() string {
	return p.validatorHRP
}

// ConsensusHRP gets the human readable prefix for an Account's consensus address.
func (p Params) ConsensusHRP() string {
	return p.consensusHRP
}

// EncodingConfig gets the encoding configuration for the network.
func (p Params) EncodingConfig() EncodingConfig {
	return p.encodingConfig
}

// EncodingConfig encapsulates the encoding config of a network.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig creates a new EncodingConfig.
func MakeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

// Supported Networks
const (
	// Cosmos identifier.
	Cosmos string = "cosmos"
	// Umee identifier.
	Umee string = "umee"
	// Osmosis identifier
	Osmosis string = "osmosis"
)
