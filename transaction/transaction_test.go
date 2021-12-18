package transaction

import (
	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/messages"
	"github.com/PineStreetLabs/nebula/networks"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestBasicTransactionFlow(t *testing.T) {
	var sender, recipient sdk.AccAddress
	var sk cryptotypes.PrivKey
	var err error

	{
		sk = account.NewPrivateKey()
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

	msg := messages.BankSend(sender, recipient, sdk.NewCoins(sdk.NewInt64Coin("atom", 10)))
	fee := sdk.NewCoins(sdk.NewInt64Coin("atom", 1))

	txn, err := Build([]sdk.Msg{msg}, 100, fee, "", 1)
	if err != nil {
		t.Fatal(err)
	}

	signerData := NewSignerData("", 0, 0)

	cfg := params.MakeTestEncodingConfig()

	_, err = Sign(cfg.TxConfig, txn, *signerData, sk)
	if err != nil {
		t.Fatal(err)
	}
}
