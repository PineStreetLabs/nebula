package networks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	umee "github.com/umee-network/umee/app"
)

func GetUmeeCfg() *sdk.Config {
	cfg := sdk.NewConfig()
	cfg.SetBech32PrefixForAccount(umee.AccountAddressPrefix, umee.AccountPubKeyPrefix)
	return cfg
}
