package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/keychain"
	"github.com/PineStreetLabs/nebula/messages"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/transaction"
	"github.com/PineStreetLabs/nebula/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
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

	var sk *secp256k1.PrivKey

	switch {
	case ctx.IsSet("from_sk"):
		seed := ctx.String("from_sk")

		sk, err = account.PrivateKeyFromHex(seed)
		if err != nil {
			return err
		}
	case ctx.IsSet("from_mnemonic"):
		mnemonic := ctx.String("from_mnemonic")

		master, err := keychain.FromMnemonic(mnemonic, "")
		if err != nil {
			return err
		}

		path, err := hd.NewParamsFromPath("m/44'/118'/0'/0/0")
		if err != nil {
			return err
		}

		key, err := keychain.Derive(master, path)
		if err != nil {
			return err
		}

		sk = &secp256k1.PrivKey{Key: key}
	default:
		sk = account.NewPrivateKey()
	}

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

	recipientAcc, err := account.AddressFromString(ctx.String("recipient"))
	if err != nil {
		return err
	}

	fmt.Println("from: " + acc.GetAddress().String())
	fmt.Println("to: " + recipientAcc.String())

	msg := messages.BankSend(acc.GetAddress(), recipientAcc, utils.NewCoin(cfg, sdk.NewIntFromUint64(ctx.Uint64("amount"))))
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
