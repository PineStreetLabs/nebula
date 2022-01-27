package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

func SlashingUnjail(address sdk.Address) slashingtypes.MsgUnjail {
	return slashingtypes.MsgUnjail{
		ValidatorAddr: address.String(),
	}
}
