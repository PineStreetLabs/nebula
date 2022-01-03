package account

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/PineStreetLabs/nebula/networks"
)

func TestFromPublicKey(t *testing.T) {
	pk, err := ParsePublicKey("AqKNXMp4eXSWIpsa/QWpNnyOCczNhKCOE/XohdspkpfI")
	if err != nil {
		t.Fatal(err)
	}

	// umee
	{
		acc, err := FromPublicKey(networks.GetUmeeCfg(), pk, 2, 1)
		if err != nil {
			t.Fatal(err)
		}

		if acc.GetAddress().String() != "umee1gfaks828le44whkqwyxwp92rg5ewt0qaucuhq8" {
			t.Fatalf("got %s\n", acc.GetAddress().String())
		}

		if acc.sequence != 1 {
			t.Fatalf("got %d : want 1", acc.sequence)
		}

		if acc.number != 2 {
			t.Fatalf("got %d : want 2", acc.number)
		}
	}

	// cosmos
	{
		acc, err := FromPublicKey(networks.GetCosmosCfg(), pk, 0, 1)
		if err != nil {
			t.Fatal(err)
		}

		if acc.GetAddress().String() != "cosmos1gfaks828le44whkqwyxwp92rg5ewt0qawwpgy4" {
			t.Fatalf("got %s\n", acc.GetAddress().String())
		}
	}
}

func TestPrivateKeyFromHex(t *testing.T) {
	sk := NewPrivateKey()
	buf := hex.EncodeToString(sk.Bytes())
	k, err := PrivateKeyFromHex(buf)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(k.Bytes(), sk.Bytes()) {
		t.Fatal("private keys do not equal")
	}
}

func TestLengthLimits(t *testing.T) {
	for i := 0; i < 100; i++ {
		sk := NewPrivateKey()
		if keyLen := len(sk.PubKey().Bytes()); keyLen != 33 {
			t.Fatalf("public key length : expected 33 : got %d", keyLen)
		}

		if addrLen := len(sk.PubKey().Address().Bytes()); addrLen != 20 {
			t.Fatalf("address key length : expected 20 : got %d", addrLen)
		}
	}
}
