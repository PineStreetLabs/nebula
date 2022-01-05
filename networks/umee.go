package networks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	umee "github.com/umee-network/umee/app"
	"github.com/umee-network/umee/app/params"
)

func GetUmeeCfg() *Params {
	return &Params{
		denom:               umee.BondDenom,
		accountHRP:          umee.AccountAddressPrefix,
		validatorHRP:        umee.ValidatorAddressPrefix,
		consensusHRP:        umee.ConsNodeAddressPrefix,
		VerifyAddressFormat: umee.VerifyAddressFormat,
		encodingConfig:      EncodingConfig(params.MakeEncodingConfig()),
	}
}

func GetCosmosCfg() *Params {
	return &Params{
		denom:               sdk.DefaultBondDenom,
		accountHRP:          sdk.Bech32PrefixAccAddr,
		validatorHRP:        sdk.Bech32PrefixValAddr,
		consensusHRP:        sdk.Bech32PrefixConsAddr,
		VerifyAddressFormat: sdk.VerifyAddressFormat,
	}
}
