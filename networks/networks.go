package networks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	umee "github.com/umee-network/umee/app"
	"github.com/umee-network/umee/app/params"
)

func GetUmeeCfg() *Params {
	modules := umee.ModuleBasics
	encCfg := params.MakeEncodingConfig()
	modules.RegisterInterfaces(encCfg.InterfaceRegistry)
	modules.RegisterLegacyAminoCodec(encCfg.Amino)

	return &Params{
		denom:               umee.BondDenom,
		accountHRP:          umee.AccountAddressPrefix,
		validatorHRP:        umee.ValidatorAddressPrefix,
		consensusHRP:        umee.ConsNodeAddressPrefix,
		VerifyAddressFormat: umee.VerifyAddressFormat,
		encodingConfig: EncodingConfig{
			InterfaceRegistry: encCfg.InterfaceRegistry,
			Marshaler:         encCfg.Marshaler,
			TxConfig:          encCfg.TxConfig,
			Amino:             encCfg.Amino,
		},
	}
}

func GetCosmosCfg() *Params {
	modules := module.NewBasicManager(
		bank.AppModuleBasic{},
		authz.AppModuleBasic{},
		auth.AppModuleBasic{},
	)
	encCfg := MakeEncodingConfig()
	modules.RegisterInterfaces(encCfg.InterfaceRegistry)
	modules.RegisterLegacyAminoCodec(encCfg.Amino)

	return &Params{
		denom:               sdk.DefaultBondDenom,
		accountHRP:          sdk.Bech32PrefixAccAddr,
		validatorHRP:        sdk.Bech32PrefixValAddr,
		consensusHRP:        sdk.Bech32PrefixConsAddr,
		VerifyAddressFormat: sdk.VerifyAddressFormat,
		encodingConfig:      encCfg,
	}
}
