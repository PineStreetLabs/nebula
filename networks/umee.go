package networks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	umee "github.com/umee-network/umee/app"
)

func GetUmeeCfg() *Params {
	return &Params{
		accountHRP:          umee.AccountAddressPrefix,
		validatorHRP:        umee.ValidatorAddressPrefix,
		consensusHRP:        umee.ConsNodeAddressPrefix,
		VerifyAddressFormat: umee.VerifyAddressFormat,
	}
	//cfg := sdk.NewConfig()
	//cfg.SetBech32PrefixForAccount(umee.AccountAddressPrefix, umee.AccountPubKeyPrefix)
	//cfg.SetBech32PrefixForValidator(umee.ValidatorAddressPrefix, umee.ValidatorPubKeyPrefix)
	//cfg.SetBech32PrefixForConsensusNode(umee.ConsNodeAddressPrefix, umee.ConsNodePubKeyPrefix)
	//cfg.SetAddressVerifier(umee.VerifyAddressFormat)
}

func GetCosmosCfg() *Params {
	return &Params{
		accountHRP:          sdk.Bech32PrefixAccAddr,
		validatorHRP:        sdk.Bech32PrefixValAddr,
		consensusHRP:        sdk.Bech32PrefixConsAddr,
		VerifyAddressFormat: sdk.VerifyAddressFormat,
	}
}
