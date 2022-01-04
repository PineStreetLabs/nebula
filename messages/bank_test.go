package messages

import (
	"testing"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestBankSend(t *testing.T) {
	var sender, recipient *account.Account
	var err error

	{
		sk := account.NewPrivateKey()
		sender, err = account.FromPublicKey(networks.GetCosmosCfg(), sk.PubKey(), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		sk := account.NewPrivateKey()
		recipient, err = account.FromPublicKey(networks.GetCosmosCfg(), sk.PubKey(), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
	}

	msg := BankSend(sender.GetAddress(), recipient.GetAddress(), sdk.NewCoins(sdk.NewInt64Coin("atom", 10)))
	t.Log(msg.String())
}
