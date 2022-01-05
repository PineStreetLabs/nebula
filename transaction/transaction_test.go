package transaction

import (
	"testing"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/messages"
	"github.com/PineStreetLabs/nebula/networks"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestBasicTransactionFlow(t *testing.T) {
	var sender, recipient *account.Account
	var sk cryptotypes.PrivKey
	var err error

	{
		sk = account.NewPrivateKey()
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

	msg := messages.BankSend(sender.GetAddress(), recipient.GetAddress(), sdk.NewCoins(sdk.NewInt64Coin("atom", 10)))
	fee := sdk.NewCoins(sdk.NewInt64Coin("atom", 1))

	txn, err := Build(networks.GetUmeeCfg(), []sdk.Msg{msg}, 100, fee, "", 1, nil)
	if err != nil {
		t.Fatal(err)
	}

	signerData := NewSignerData("", 0, 0)

	signedTxn, err := Sign(networks.GetUmeeCfg().EncodingConfig().TxConfig, txn, *signerData, sk)
	if err != nil {
		t.Fatal(err)
	}

	res, err := Serialize(networks.GetUmeeCfg().EncodingConfig().TxConfig, signedTxn)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%x\n", res)
}
