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

// BankSend returns a MsgSend message.
func BankSend(from, to sdk.Address, coins ...sdk.Coin) *banktypes.MsgSend {
	return &banktypes.MsgSend{
		FromAddress: from.String(),
		ToAddress:   to.String(),
		Amount:      coins,
	}
}

// NewInput is a helper function to create inputs.
func NewInput(addr sdk.Address, coins ...sdk.Coin) banktypes.Input {
	return banktypes.Input{
		Address: addr.String(),
		Coins:   coins,
	}
}

// NewOutput is a helper function to create outputs.
func NewOutput(addr sdk.Address, coins ...sdk.Coin) banktypes.Output {
	return banktypes.Output{
		Address: addr.String(),
		Coins:   coins,
	}
}

// BankMultiSend returns a MsgMultiSend message.
func BankMultiSend(ins []banktypes.Input, outs []banktypes.Output) *banktypes.MsgMultiSend {
	return &banktypes.MsgMultiSend{
		Inputs:  ins,
		Outputs: outs,
	}
}
