package messages

import (
	"encoding/base64"
	"math/big"
	"testing"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/utils"
	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestStaking(t *testing.T) {
	key, err := base64.StdEncoding.DecodeString("AxbeWcFto+wFfoQOot6oPJpeqf5j3sbl6hiMhYhb+ON7")
	if err != nil {
		t.Fatal(err)
	}

	pk, err := account.Secp256k1PublicKey(key)
	if err != nil {
		panic(err)
	}
	validator, err := account.NewValidatorAccount(networks.GetCosmosCfg(), pk, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	valAddress, err := account.ToValidatorAddress(networks.GetCosmosCfg(), validator)
	if err != nil {
		t.Fatal(err)
	}

	msg, err := StakingCreateValidator(valAddress, pk, utils.NewCoinFromBigInt(networks.GetCosmosCfg().Denom(), big.NewInt(1)), stakingtypes.Description{}, stakingtypes.CommissionRates{}, types.Int{})
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"type":"cosmos-sdk/MsgCreateValidator","value":{"commission":{"max_change_rate":"0","max_rate":"0","rate":"0"},"delegator_address":"umee1d9aegvzrt9qz7w4n52dffsrrfvnm002050hzar","description":{},"min_self_delegation":"0","pubkey":{"type":"tendermint/PubKeySecp256k1","value":"AxbeWcFto+wFfoQOot6oPJpeqf5j3sbl6hiMhYhb+ON7"},"validator_address":"umeevaloper1d9aegvzrt9qz7w4n52dffsrrfvnm00205tsdvf","value":{"amount":"1","denom":"atom"}}}`
	if expected != string(msg.GetSignBytes()) {
		t.Fatalf("want %s, got %s", expected, msg.GetSignBytes())
	}
}
