package messages

import (
	"math/big"
	"testing"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/utils"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func TestStaking(t *testing.T) {
	var validator, validator *account.Account
	var valPK cryptotypes.PubKey
	var err error

	{
		sk := account.NewPrivateKey()
		valPK = sk.PubKey()
		validator, err = account.ValidatorFromPublicKey(networks.GetCosmosCfg(), valPK, 0, 0)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		sk := account.NewPrivateKey()
		validator, err = account.FromPublicKey(networks.GetCosmosCfg(), sk.PubKey(), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
	}

	StakingCreateValidator(validator, valPK, utils.NewCoin(networks.GetCosmosCfg(), big.NewInt(1)), nil, nil, nil)
}
