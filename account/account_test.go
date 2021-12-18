package account

import (
	"bytes"
	"encoding/hex"
	"github.com/PineStreetLabs/nebula/networks"
	"testing"
)

func TestFromPublicKey(t *testing.T) {
	sk := NewPrivateKey()
	acc, err := FromPublicKey(networks.GetUmeeCfg(), sk.PubKey())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", acc.GetAddress().String())
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
