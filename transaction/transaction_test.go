package transaction

import (
	"bytes"
	"encoding/hex"
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
		sk, err = account.PrivateKeyFromHex("173200b80396709fe10b3a44df5982ce0834a2d6fb1e57ca55efe6ebd290d776")
		if err != nil {
			t.Fatal(err)
		}

		sender, err = account.NewUserAccount(networks.GetCosmosCfg(), sk.PubKey(), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		sk, err = account.PrivateKeyFromHex("9e6dfe0d36d7cdbef611e05157e4c5e4bd161221238136a35be177e614b748b3")
		if err != nil {
			t.Fatal(err)
		}
		recipient, err = account.NewUserAccount(networks.GetCosmosCfg(), sk.PubKey(), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
	}

	msg := messages.BankSend(sender.GetAddress(), recipient.GetAddress(), sdk.NewCoins(sdk.NewInt64Coin("atom", 10))...)
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

	expected, _ := hex.DecodeString("0a8f010a8a010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e64126a0a2d636f736d6f7331636d716873386a687468666d7078757561326a6d616d746b78746477376d30656d32366d6e6d122d636f736d6f7331346c66716879676672386c61666c766b74636c71646d37653661666668757a35326c616430661a0a0a0461746f6d120231301801125f0a4e0a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a210345f45bdfd75c5070843273583831485aa434ed7bac0c95b5bed9f2673f80fa6512040a020801120d0a090a0461746f6d12013110641a40a8b201d014bb3e35cc101c4074af36cb6ee78379b52529604a4ca7e917d3df182ec1270c28064f54d061dfc543853e78dbb32df97e5f75c56fb2561d38cb0e18")
	if !bytes.Equal(res, expected) {
		t.Fatalf("transaction bytes not expected\n%x\n%x", res, expected)
	}
}
