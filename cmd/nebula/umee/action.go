package umee

import (
	"fmt"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/cmd/nebula/common"
	"github.com/PineStreetLabs/nebula/messages/umee"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/transaction"
	"github.com/PineStreetLabs/nebula/utils"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/urfave/cli"
)

func lendAsset(ctx *cli.Context) (err error) {
	cfg, acc, sk, err := getAccountAndNetworkConf(ctx)
	if err != nil {
		return err
	}

	msg := umee.NewMsgLendAsset(acc.GetAddress(), utils.NewCoinFromUint64(cfg.Denom(), ctx.Uint64("amount")))

	return buildAndSignTx(ctx, msg, acc, sk)
}

func withdrawAsset(ctx *cli.Context) (err error) {
	cfg, acc, sk, err := getAccountAndNetworkConf(ctx)
	if err != nil {
		return err
	}

	msg := umee.NewMsgWithdrawAsset(acc.GetAddress(), utils.NewCoinFromUint64(cfg.Denom(), ctx.Uint64("amount")))

	return buildAndSignTx(ctx, msg, acc, sk)
}

func setCollateral(ctx *cli.Context) (err error) {
	cfg, acc, sk, err := getAccountAndNetworkConf(ctx)
	if err != nil {
		return err
	}

	msg := umee.NewMsgSetCollateral(acc.GetAddress(), cfg.Denom(), ctx.Bool("enabled"))

	return buildAndSignTx(ctx, msg, acc, sk)
}

func repayAsset(ctx *cli.Context) (err error) {
	cfg, acc, sk, err := getAccountAndNetworkConf(ctx)
	if err != nil {
		return err
	}

	msg := umee.NewMsgRepayAsset(acc.GetAddress(), utils.NewCoinFromUint64(cfg.Denom(), ctx.Uint64("amount")))

	return buildAndSignTx(ctx, msg, acc, sk)
}

func getAccountAndNetworkConf(ctx *cli.Context) (*networks.Params, *account.Account, cryptotypes.PrivKey, error) {
	cfg, err := common.GetNetworkConfig(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	sk, err := account.PrivateKeyFromHex(ctx.String("private_key"))
	if err != nil {
		return nil, nil, nil, err
	}

	acc, err := account.NewUserAccount(cfg, sk.PubKey(), ctx.Uint64("acc_number"), ctx.Uint64("acc_sequence"))
	if err != nil {
		return nil, nil, nil, err
	}

	return cfg, acc, sk, nil
}

func buildAndSignTx(ctx *cli.Context, msg sdk.Msg, acc *account.Account, sk cryptotypes.PrivKey) (err error) {
	cfg, err := common.GetNetworkConfig(ctx)
	if err != nil {
		return err
	}

	gasLimit := ctx.Uint64("gas_limit")
	fee := sdk.NewCoins(sdk.NewInt64Coin("uumee", ctx.Int64("fee")))
	timeoutHeight := ctx.Uint64("timeout_height")
	memo := ctx.String("memo")

	txnBuilder, err := transaction.Build(cfg, []sdk.Msg{msg}, gasLimit, fee, memo, timeoutHeight)
	if err != nil {
		return err
	}

	signerData := transaction.NewSignerData(ctx.String("chain_id"), acc.GetAccountNumber(), acc.GetSequence())
	txn, err := transaction.Sign(cfg.EncodingConfig().TxConfig, txnBuilder, *signerData, sk)
	if err != nil {
		return err
	}

	serializedTxn, err := transaction.Serialize(cfg.EncodingConfig().TxConfig, txn)
	if err != nil {
		return err
	}

	fmt.Printf("%x\n", serializedTxn)
	return nil
}
