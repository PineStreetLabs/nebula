package networks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	umee "github.com/umee-network/umee/app"
)

// Params includes the scope and configuration of each network.
// In the cosmos-sdk, this is akin to the sdk.Config type.
// We choose not to use sdk.Config because of its use of init and how it is used globally in the cosmos-sdk library.
// Params tightly scopes relevant app chain configurations to work neatly across this library.
type Params struct {
	accountHRP   string
	validatorHRP string
	consensusHRP string
}

func (p Params) AccountHRP() string {
	return p.accountHRP
}

func (p Params) ValidatorHRP() string {
	return p.validatorHRP
}

func (p Params) ConsensusHRP() string {
	return p.consensusHRP
}

func GetUmeeCfg() *Params {
	return &Params{
		accountHRP:   umee.AccountAddressPrefix,
		validatorHRP: umee.ValidatorAddressPrefix,
		consensusHRP: umee.ConsNodeAddressPrefix,
	}
	//cfg := sdk.NewConfig()
	//cfg.SetBech32PrefixForAccount(umee.AccountAddressPrefix, umee.AccountPubKeyPrefix)
	//cfg.SetBech32PrefixForValidator(umee.ValidatorAddressPrefix, umee.ValidatorPubKeyPrefix)
	//cfg.SetBech32PrefixForConsensusNode(umee.ConsNodeAddressPrefix, umee.ConsNodePubKeyPrefix)
	//cfg.SetAddressVerifier(umee.VerifyAddressFormat)
}

func GetCosmosCfg() *Params {
	return &Params{
		accountHRP:   sdk.Bech32PrefixAccAddr,
		validatorHRP: sdk.Bech32PrefixValAddr,
		consensusHRP: sdk.Bech32PrefixConsAddr,
	}
}
