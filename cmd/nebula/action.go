package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/messages"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/transaction"
	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
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
	acc, err := account.FromPublicKey(cfg, sk.PubKey(), 0, 0)
	if err != nil {
		return err
	}

	type Result struct {
		Address    string `json:"address"`
		PrivateKey string `json:"private_key"`
	}

	resp, err := json.Marshal(Result{
		Address:    acc.GetAddress().String(),
		PrivateKey: hex.EncodeToString(sk.Bytes()),
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

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

	acc, err := account.FromPublicKey(cfg, sk.PubKey(), ctx.Uint64("acc_number"), ctx.Uint64("acc_sequence"))
	if err != nil {
		return err
	}

	recipientAcc, err := account.FromAddress(cfg, ctx.String("recipient"))
	if err != nil {
		return err
	}

	fmt.Println("from: " + acc.GetAddress().String())
	fmt.Println("to: " + recipientAcc.GetAddress().String())

	msg := messages.BankSend(acc.GetAddress(), recipientAcc.GetAddress(), sdk.NewCoins(sdk.NewInt64Coin("uumee", 1000)))
	gasLimit := ctx.Uint64("gas_limit")
	fee := sdk.NewCoins(sdk.NewInt64Coin("uumee", ctx.Int64("fee")))
	timeoutHeight := ctx.Uint64("timeout_height")
	memo := ctx.String("memo")

	txnBuilder, err := transaction.Build(cfg, []sdk.Msg{msg}, gasLimit, fee, memo, timeoutHeight, []cryptotypes.PubKey{acc.GetPubKey()})
	if err != nil {
		return err
	}

	signerData := transaction.NewSignerData("umee-local-testnet", ctx.Uint64("acc_number"), ctx.Uint64("acc_sequence"))
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
	rpcClient, err := http.New("http://0.0.0.0:26657", "/")
	if err != nil {
		return err
	}

	response, err := client.TxServiceBroadcast(context.Background(), client.Context{Client: rpcClient}, &broadcastTxRequest)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", response.TxResponse)
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
