package messages

import (
	"testing"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/utils"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func TestBankSend(t *testing.T) {
	sender, err := account.AddressFromString("cosmos1d9h8qat57ljhcm")
	if err != nil {
		t.Fatal(err)
	}

	recipient, err := account.AddressFromString("cosmos1da6hgur4wsmpnjyg")
	if err != nil {
		t.Fatal(err)
	}

	msg := BankSend(sender, recipient, utils.NewCoinFromUint64(networks.GetCosmosCfg().Denom(), 10))
	expected := `{"type":"cosmos-sdk/MsgSend","value":{"amount":[{"amount":"10","denom":"atom"}],"from_address":"cosmos1d9h8qat57ljhcm","to_address":"cosmos1da6hgur4wsmpnjyg"}}`
	if expected != string(msg.GetSignBytes()) {
		t.Fatalf("want %s, got %s", expected, msg.GetSignBytes())
	}
}

func TestBankMultiSend(t *testing.T) {
	sender, err := account.AddressFromString("cosmos1d9h8qat57ljhcm")
	if err != nil {
		t.Fatal(err)
	}

	recipient, err := account.AddressFromString("cosmos1da6hgur4wsmpnjyg")
	if err != nil {
		t.Fatal(err)
	}

	ins := make([]banktypes.Input, 0)
	ins = append(ins, NewInput(sender, utils.NewCoinFromUint64(networks.GetCosmosCfg().Denom(), 10)))

	outs := make([]banktypes.Output, 0)
	outs = append(outs, NewOutput(recipient, utils.NewCoinFromUint64(networks.GetCosmosCfg().Denom(), 10)))

	msg := BankMultiSend(ins, outs)
	expected := `{"type":"cosmos-sdk/MsgMultiSend","value":{"inputs":[{"address":"cosmos1d9h8qat57ljhcm","coins":[{"amount":"10","denom":"atom"}]}],"outputs":[{"address":"cosmos1da6hgur4wsmpnjyg","coins":[{"amount":"10","denom":"atom"}]}]}}`
	if expected != string(msg.GetSignBytes()) {
		t.Fatalf("want %s, got %s", expected, msg.GetSignBytes())
	}
}
