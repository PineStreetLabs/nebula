package umee

import sdk "github.com/cosmos/cosmos-sdk/types"
import umeetypes "github.com/umee-network/umee/x/leverage/types"

func NewMsgLendAsset(lender sdk.Address, amount sdk.Coin) *umeetypes.MsgLendAsset {
	return &umeetypes.MsgLendAsset{
		Lender: lender.String(),
		Amount: amount,
	}
}

func NewMsgWithdrawAsset(lender sdk.Address, amount sdk.Coin) *umeetypes.MsgWithdrawAsset {
	return &umeetypes.MsgWithdrawAsset{
		Lender: lender.String(),
		Amount: amount,
	}
}

func NewMsgSetCollateral(borrower sdk.Address, denom string, enable bool) *umeetypes.MsgSetCollateral {
	return &umeetypes.MsgSetCollateral{
		Borrower: borrower.String(),
		Denom:    denom,
		Enable:   enable,
	}
}

func NewMsgRepayAsset(borrower sdk.Address, amount sdk.Coin) *umeetypes.MsgRepayAsset {
	return &umeetypes.MsgRepayAsset{
		Borrower: borrower.String(),
		Amount:   amount,
	}
}
