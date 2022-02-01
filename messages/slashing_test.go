package messages

import (
	"testing"

	"github.com/PineStreetLabs/nebula/account"
)

func TestSlashingUnjail(t *testing.T) {
	addr, err := account.AddressFromString("cosmosvaloper1v93xxeqhg9nn6")
	if err != nil {
		t.Fatal(err)
	}

	msg := SlashingUnjail(addr)
	expected := `{"type":"cosmos-sdk/MsgUnjail","value":{"address":"cosmosvaloper1v93xxeqhg9nn6"}}`
	if expected != string(msg.GetSignBytes()) {
		t.Fatalf("want %s, got %s", expected, msg.GetSignBytes())
	}
}
