package account

import (
	"encoding/hex"
	"testing"

	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/crypto/types"
)

func TestNewMultiSig(t *testing.T) {
	pk0, _ := hex.DecodeString("02452611abd6595aefec1889a0244c28ebeb78e1fa490e1d61f6af1f3d7722899d")
	pk, _ := Secp256k1PublicKey(pk0)

	keys := []types.PubKey{
		pk,
		pk,
	}

	acc, err := NewMultiSigAccount(networks.GetUmeeCfg(), 2, keys, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	expected := "umee1y2jtwhxu82t5n6wl5uqjqwwcpg8et4drkr2hvn"
	if addr := acc.GetAddress().String(); addr != expected {
		t.Fatalf("unexpected address : got %s, wanted %s", addr, expected)
	}
}
