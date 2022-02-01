package messages

import (
	"testing"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func TestGovDeposit(t *testing.T) {
	depositor, err := account.AddressFromString("cosmos1v9jxgu33kfsgr5")
	if err != nil {
		t.Fatal(err)
	}

	msg := GovDeposit(0, depositor, utils.NewCoinFromUint64("stake", 1000))
	expected := `{"type":"cosmos-sdk/MsgDeposit","value":{"amount":[{"amount":"1000","denom":"stake"}],"depositor":"cosmos1v9jxgu33kfsgr5","proposal_id":"0"}}`
	if expected != string(msg.GetSignBytes()) {
		t.Fatalf("want %s, got %s", expected, msg.GetSignBytes())
	}
}

func TestGovSubmitProposal(t *testing.T) {
	proposer, err := account.AddressFromString("cosmos1v9jxgu33kfsgr5")
	if err != nil {
		t.Fatal(err)
	}

	coins := make([]sdk.Coin, 0)
	msg := GovSubmitProposal(govtypes.NewTextProposal("test", "abcd"), proposer, coins...)
	expected := `{"type":"cosmos-sdk/MsgSubmitProposal","value":{"content":{"type":"cosmos-sdk/TextProposal","value":{"description":"abcd","title":"test"}},"initial_deposit":[],"proposer":"cosmos1v9jxgu33kfsgr5"}}`
	if expected != string(msg.GetSignBytes()) {
		t.Fatalf("want %s, got %s", expected, msg.GetSignBytes())
	}
}
