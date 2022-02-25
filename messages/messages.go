package messages

import (
	"github.com/PineStreetLabs/nebula/networks"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Marshal marshals the cosmos-sdk Msg interface.
func Marshal(cfg networks.EncodingConfig, msg sdk.Msg) ([]byte, error) {
	return cfg.Marshaler.MarshalInterface(msg)
}

// Unmarshal unmarshals a byte slice into a cosmos-sdk Msg interface.
func Unmarshal(cfg networks.EncodingConfig, buf []byte) (sdk.Msg, error) {
	var msg sdk.Msg
	if err := cfg.Marshaler.UnmarshalInterface(buf, &msg); err != nil {
		return nil, err
	}
	return msg, nil
}
