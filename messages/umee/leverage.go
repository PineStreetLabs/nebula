package umee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	umeetypes "github.com/umee-network/umee/x/leverage/types"
)

// NewMsgLendAsset returns a MsgLendAsset message.
func NewMsgLendAsset(lender sdk.Address, amount sdk.Coin) *umeetypes.MsgLendAsset {
	return &umeetypes.MsgLendAsset{
		Lender: lender.String(),
		Amount: amount,
	}
}

// NewMsgWithdrawAsset returns a MsgWithdrawAsset message.
func NewMsgWithdrawAsset(lender sdk.Address, amount sdk.Coin) *umeetypes.MsgWithdrawAsset {
	return &umeetypes.MsgWithdrawAsset{
		Lender: lender.String(),
		Amount: amount,
	}
}

// NewMsgSetCollateral returns a MsgSetCollateral message.
func NewMsgSetCollateral(borrower sdk.Address, denom string, enable bool) *umeetypes.MsgSetCollateral {
	return &umeetypes.MsgSetCollateral{
		Borrower: borrower.String(),
		Denom:    denom,
		Enable:   enable,
	}
}

// NewMsgRepayAsset returns a MsgRepayAsset message.
func NewMsgRepayAsset(borrower sdk.Address, amount sdk.Coin) *umeetypes.MsgRepayAsset {
	return &umeetypes.MsgRepayAsset{
		Borrower: borrower.String(),
		Amount:   amount,
	}
}
