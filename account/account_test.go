package account

import (
	"github.com/PineStreetLabs/nebula/networks"
	"testing"
)

func TestFromPublicKey(t *testing.T) {
	sk := NewPrivateKey()
	addr, err := FromPublicKey(networks.GetUmeeCfg(), sk.PubKey())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", addr.String())
}
