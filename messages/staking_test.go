package messages

import (
	"math/big"
	"testing"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/utils"
	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestStaking(t *testing.T) {
	sk := account.NewPrivateKey()
	valPK := sk.PubKey()
	validator, err := account.NewValidatorAccount(networks.GetCosmosCfg(), valPK, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	valAddress, err := account.ToValidatorAddress(networks.GetCosmosCfg(), validator)
	if err != nil {
		t.Fatal(err)
	}

	StakingCreateValidator(valAddress, valPK, utils.NewCoinFromBigInt(networks.GetCosmosCfg(), big.NewInt(1)), stakingtypes.Description{}, stakingtypes.CommissionRates{}, types.Int{})
}
