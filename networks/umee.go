package networks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	umee "github.com/umee-network/umee/app"
)

func GetUmeeCfg() *sdk.Config {
	cfg := sdk.NewConfig()
	cfg.SetBech32PrefixForAccount(umee.AccountAddressPrefix, umee.AccountPubKeyPrefix)
	cfg.SetBech32PrefixForValidator(umee.ValidatorAddressPrefix, umee.ValidatorPubKeyPrefix)
	cfg.SetBech32PrefixForConsensusNode(umee.ConsNodeAddressPrefix, umee.ConsNodePubKeyPrefix)
	cfg.SetAddressVerifier(umee.VerifyAddressFormat)
	return cfg
}

func GetCosmosCfg() *sdk.Config {
	return sdk.NewConfig()
}
