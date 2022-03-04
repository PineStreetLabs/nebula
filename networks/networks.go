package networks

import (
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"

	umee "github.com/umee-network/umee/app"
	leverageTypes "github.com/umee-network/umee/x/leverage/types"
)

// GetUmeeCfg returns the network parameters for the Umee network.
func GetUmeeCfg() *Params {
	modules := umee.ModuleBasics
	encCfg := MakeEncodingConfig()
	modules.RegisterInterfaces(encCfg.InterfaceRegistry)
	modules.RegisterLegacyAminoCodec(encCfg.Amino)
	leverageTypes.RegisterInterfaces(encCfg.InterfaceRegistry)
	cryptocodec.RegisterInterfaces(encCfg.InterfaceRegistry)

	return &Params{
		denom:               umee.BondDenom,
		accountHRP:          umee.AccountAddressPrefix,
		validatorHRP:        umee.ValidatorAddressPrefix,
		consensusHRP:        umee.ConsNodeAddressPrefix,
		VerifyAddressFormat: umee.VerifyAddressFormat,
		encodingConfig:      encCfg,
	}
}

// GetCosmosCfg returns the network parameters for the Cosmos Hub network.
func GetCosmosCfg() *Params {
	modules := module.NewBasicManager(
		bank.AppModuleBasic{},
		authz.AppModuleBasic{},
		auth.AppModuleBasic{},
	)

	encCfg := MakeEncodingConfig()
	modules.RegisterLegacyAminoCodec(encCfg.Amino)
	modules.RegisterInterfaces(encCfg.InterfaceRegistry)
	cryptocodec.RegisterInterfaces(encCfg.InterfaceRegistry)

	return &Params{
		denom:               "atom",
		accountHRP:          sdk.Bech32PrefixAccAddr,
		validatorHRP:        sdk.Bech32PrefixValAddr,
		consensusHRP:        sdk.Bech32PrefixConsAddr,
		VerifyAddressFormat: sdk.VerifyAddressFormat,
		encodingConfig:      encCfg,
	}
}
