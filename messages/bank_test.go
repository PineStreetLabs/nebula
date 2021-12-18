package messages

import (
	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestBankSend(t *testing.T) {
	var sender, recipient sdk.AccAddress
	var err error

	{
		sk := account.NewPrivateKey()
		sender, err = account.FromPublicKey(networks.GetCosmosCfg(), sk.PubKey())
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		sk := account.NewPrivateKey()
		recipient, err = account.FromPublicKey(networks.GetCosmosCfg(), sk.PubKey())
		if err != nil {
			t.Fatal(err)
		}
	}

	msg := BankSend(sender, recipient, sdk.NewCoins(sdk.NewInt64Coin("atom", 10)))
	t.Log(msg.String())
}
