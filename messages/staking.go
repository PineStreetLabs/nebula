package messages

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func StakingCreateValidator(validator sdk.ValAddress, pk cryptotypes.PubKey, value sdk.Coin, desc stakingtypes.Description, commission stakingtypes.CommissionRates, minSelfDelegation sdk.Int) (*stakingtypes.MsgCreateValidator, error) {
	// pkAny, err := codectypes.NewAnyWithValue(pk)
	// if err != nil {
	// 	return nil, err
	// }

	// sdk.ValAddress(validator)

	return stakingtypes.NewMsgCreateValidator(validator, pk, value, desc, commission, minSelfDelegation)
	// return &stakingtypes.MsgCreateValidator{
	// 	Description:       desc,
	// 	DelegatorAddress:  validator.String(),
	// 	ValidatorAddress:  validator.String(),
	// 	Pubkey:            pkAny,
	// 	Value:             value,
	// 	Commission:        commission,
	// 	MinSelfDelegation: minSelfDelegation,
	// }, nil
}

func StakingEditValidator(validator sdk.Address, desc stakingtypes.Description, rate sdk.Dec, minSelfDelegation sdk.Int) stakingtypes.MsgEditValidator {
	return stakingtypes.MsgEditValidator{
		Description:       desc,
		CommissionRate:    &rate,
		ValidatorAddress:  validator.String(),
		MinSelfDelegation: &minSelfDelegation,
	}
}

func StakingDelegate(delegate sdk.Address, validator sdk.Address, amount sdk.Coin) stakingtypes.MsgDelegate {
	return stakingtypes.MsgDelegate{
		DelegatorAddress: delegate.String(),
		ValidatorAddress: validator.String(),
		Amount:           amount,
	}
}

func StakingUndelegate(delegate sdk.Address, validator sdk.Address, amount sdk.Coin) stakingtypes.MsgUndelegate {
	return stakingtypes.MsgUndelegate{
		DelegatorAddress: delegate.String(),
		ValidatorAddress: validator.String(),
		Amount:           amount,
	}
}

func StakingBeginRedelgate(delegate, validatorSrc, validatorDest sdk.Address, amount sdk.Coin) stakingtypes.MsgBeginRedelegate {
	return stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    delegate.String(),
		ValidatorSrcAddress: validatorSrc.String(),
		ValidatorDstAddress: validatorDest.String(),
		Amount:              amount,
	}
}
