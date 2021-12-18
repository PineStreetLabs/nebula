package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

/*
Bank.go implements the x/bank functionality.
It is common for app chains to build wrappers around bank.go,
therefore all functions are network aware.
*/

func BankSend(from, to sdk.AccAddress, coins sdk.Coins) sdk.Msg {
	return banktypes.NewMsgSend(from, to, coins)
}
