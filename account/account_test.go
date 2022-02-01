package account

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/PineStreetLabs/nebula/networks"
)

func TestFromPublicKey(t *testing.T) {
	pk, err := ParseSecp256k1PublicKey("AqKNXMp4eXSWIpsa/QWpNnyOCczNhKCOE/XohdspkpfI")
	if err != nil {
		t.Fatal(err)
	}

	// umee
	{
		acc, err := NewUserAccount(networks.GetUmeeCfg(), pk, 2, 1)
		if err != nil {
			t.Fatal(err)
		}

		if acc.GetAddress().String() != "umee1gfaks828le44whkqwyxwp92rg5ewt0qaucuhq8" {
			t.Fatalf("got %s\n", acc.GetAddress().String())
		}

		if acc.sequence != 1 {
			t.Fatalf("got %d : want 1", acc.sequence)
		}

		if acc.number != 2 {
			t.Fatalf("got %d : want 2", acc.number)
		}
	}

	// cosmos
	{
		acc, err := NewUserAccount(networks.GetCosmosCfg(), pk, 0, 1)
		if err != nil {
			t.Fatal(err)
		}

		if acc.GetAddress().String() != "cosmos1gfaks828le44whkqwyxwp92rg5ewt0qawwpgy4" {
			t.Fatalf("got %s\n", acc.GetAddress().String())
		}
	}

	// Missing public key.
	{
		_, err := fromPublicKey("hrp", networks.GetCosmosCfg(), nil, 0, 0)
		if err == nil || err != errMissingPublicKey {
			t.Fatal("expected err")
		}
	}
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

func TestUserAccount(t *testing.T) {
	pk, err := ParseSecp256k1PublicKey("Ajo3B71w4LnB6jo/r4g7MbYL6cNwd766lYyhXY9ae99M")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := NewUserAccount(networks.GetUmeeCfg(), pk, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	expectedAddr := "umee1jgsy5wnugmqjlx0cnn4cttfz68gjwk4exskzay"
	if addr := acc.GetAddress().String(); addr != expectedAddr {
		t.Fatalf("expected %s : got %s", expectedAddr, addr)
	}

	// Set the sequence.
	acc.SetSequence(2)

	if seq := acc.GetSequence(); seq != 2 {
		t.Fatalf("expected sequence 2 : got %d", seq)
	}

	// Convert to validator address. Should fail.
	if _, err := ToValidatorAddress(networks.GetUmeeCfg(), acc); err == nil {
		t.Fatal("expected err")
	}
}

func TestValAccount(t *testing.T) {
	pk, err := ParseSecp256k1PublicKey("Ajo3B71w4LnB6jo/r4g7MbYL6cNwd766lYyhXY9ae99M")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := NewValidatorAccount(networks.GetUmeeCfg(), pk, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	expectedAddr := "umeevaloper1jgsy5wnugmqjlx0cnn4cttfz68gjwk4ex53dvw"
	if addr := acc.GetAddress().String(); addr != expectedAddr {
		t.Fatalf("expected %s : got %s", expectedAddr, addr)
	}

	if _, err := ToValidatorAddress(networks.GetUmeeCfg(), acc); err != nil {
		t.Fatal(err)
	}
}

func TestConsAccount(t *testing.T) {
	pk, err := ParseEd25519PublicKey("fxCmUUD3ijcE+mBPs4JAjvl5+sp3B03tTGSBXuEUfpU=")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := NewConsensusAccount(networks.GetUmeeCfg(), pk, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	expectedAddr := "umeevalcons13ttqjm8569aefdkrl4xw6tu5awlvhdt372zyxz"
	if addr := acc.GetAddress().String(); addr != expectedAddr {
		t.Fatalf("expected %s : got %s", expectedAddr, addr)
	}
}

func TestNewAccount(t *testing.T) {
	pk, err := ParseSecp256k1PublicKey("AwAOXeWgNf1FjMaayrSnrOOKz+Fivr6DiI/i0x0sZCHw")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := NewAccount("cosmos14pt0q5cwf38zt08uu0n6yrstf3rndzr5057jys", pk, 10, 1); err != nil {
		t.Fatal(err)
	}

	if _, err := NewAccount("cosmos14pt0q5cwf38zt08uu0n6yrstf3rndzr5057jys", nil, 10, 1); err == nil || err != errMissingPublicKey {
		t.Fatalf("expected %v", errMissingPublicKey)
	}

}

func TestMarshal(t *testing.T) {
	pk, err := ParseSecp256k1PublicKey("AwAOXeWgNf1FjMaayrSnrOOKz+Fivr6DiI/i0x0sZCHw")
	if err != nil {
		t.Fatal(err)
	}
	acc, err := NewAccount("cosmos14pt0q5cwf38zt08uu0n6yrstf3rndzr5057jys", pk, 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := acc.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	newAcct := Account{}
	(&newAcct).UnmarshalJSON(buf)

	if newAcct.GetAddress().String() != "cosmos14pt0q5cwf38zt08uu0n6yrstf3rndzr5057jys" {
		t.Fatal("address does not match")
	}

	if !bytes.Equal(newAcct.GetPubKey().Bytes(), acc.GetPubKey().Bytes()) {
		t.Fatal("public key does not match")
	}

	if newAcct.GetAccountNumber() != acc.GetAccountNumber() {
		t.Fatal("account number does not match")
	}

	if newAcct.GetSequence() != acc.GetSequence() {
		t.Fatal("account sequence does not match")
	}

	// Try with null public key.
	{
		newAcct := Account{}
		acc.publicKey = nil
		buf, err := acc.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		(&newAcct).UnmarshalJSON(buf)

		if len(newAcct.GetPubKey().Bytes()) != 0 {
			t.Fatal("expected empty public key")
		}
	}
}
