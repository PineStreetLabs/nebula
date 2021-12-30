package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/transaction"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	"github.com/urfave/cli"
)

var errUnsupportedNetwork = errors.New("unsupported network")

func newAccount(ctx *cli.Context) (err error) {
	cfg, err := getNetworkConfig(ctx)
	if err != nil {
		return err
	}

	sk := account.NewPrivateKey()
	acc, err := account.FromPublicKey(cfg, sk.PubKey())
	if err != nil {
		return err
	}

	fmt.Printf("%x\n", sk.Bytes())
	fmt.Printf("%s\n", acc.GetAddress().String())
	return nil
}

func newBankSend(ctx *cli.Context) (err error) {
	cfg, err := getNetworkConfig(ctx)
	if err != nil {
		return err
	}

	sk, err := account.PrivateKeyFromHex(ctx.String("private_key"))
	if err != nil {
		return err
	}
	acc, err := account.FromPublicKey(cfg, sk.PubKey())
	if err != nil {
		return err
	}

	recipientSk, err := account.PrivateKeyFromHex("d3d0bf0a4d0a3844fdb600950c206eb455b72e487a7dc28a51c0d4332f1f62bb")
	if err != nil {
		return err
	}
	recipientAcc, err := account.FromPublicKey(cfg, recipientSk.PubKey())
	if err != nil {
		return err
	}

	fmt.Println("from: " + acc.GetAddress().String())
	fmt.Println("to: " + recipientAcc.GetAddress().String())

	msg := bankTypes.NewMsgSend(acc.GetAddress().Bytes(), recipientAcc.GetAddress().Bytes(), sdk.NewCoins(sdk.NewInt64Coin("uumee", 1000)))
	gasLimit := ctx.Uint64("gas_limit")
	fee := sdk.NewCoins(sdk.NewInt64Coin("uumee", ctx.Int64("fee")))
	timeoutHeight := ctx.Uint64("timeout_height")
	memo := ctx.String("memo")

	txnBuilder, err := transaction.Build(cfg, []sdk.Msg{msg}, gasLimit, fee, memo, timeoutHeight)
	if err != nil {
		return err
	}

	signerData := transaction.NewSignerData("chain-id", acc.GetAccountNumber(), acc.GetSequence())
	txn, err := transaction.Sign(cfg.EncodingConfig().TxConfig, txnBuilder, *signerData, sk)
	if err != nil {
		return err
	}

	serializedTxn, err := transaction.Serialize(cfg.EncodingConfig().TxConfig, txn)
	if err != nil {
		return err
	}

	broadcastTxRequest := tx.BroadcastTxRequest{
		TxBytes: serializedTxn,
		Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
	}
	rpcClient, err := http.New("0.0.0.0:9090", "/websocket")
	response, err := client.TxServiceBroadcast(context.Background(), client.Context{Client: rpcClient}, &broadcastTxRequest)
	if err != nil {
		return err
	}
	fmt.Printf("%d\n", response.TxResponse.Code)
	return nil
}

func getNetworkConfig(ctx *cli.Context) (*networks.Params, error) {
	switch network := ctx.String("network"); network {
	case networks.Cosmos:
		return networks.GetCosmosCfg(), nil
	case networks.Umee:
		return networks.GetUmeeCfg(), nil
	default:
		return nil, fmt.Errorf("%w : %s", errUnsupportedNetwork, network)
	}
}
