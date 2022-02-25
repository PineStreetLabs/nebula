package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/PineStreetLabs/nebula/cmd/nebula/common"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/keychain"
	"github.com/PineStreetLabs/nebula/messages"
	"github.com/PineStreetLabs/nebula/transaction"
	"github.com/PineStreetLabs/nebula/utils"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/urfave/cli"
)

func newAccount(ctx *cli.Context) (err error) {
	cfg, err := common.GetNetworkConfig(ctx)
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

	acc, err := account.NewUserAccount(cfg, sk.PubKey(), 0, 0)
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
	cfg, err := common.GetNetworkConfig(ctx)
	if err != nil {
		return err
	}

	senderAcc, err := account.AddressFromString(ctx.String("sender"))
	if err != nil {
		return err
	}

	recipientAcc, err := account.AddressFromString(ctx.String("recipient"))
	if err != nil {
		return err
	}

	msg := messages.BankSend(senderAcc, recipientAcc, utils.NewCoinFromUint64(cfg.Denom(), ctx.Uint64("amount")))

	serializedMsg, err := messages.Marshal(cfg.EncodingConfig(), msg)
	if err != nil {
		return err
	}

	fmt.Printf("%x\n", serializedMsg)
	return nil
}

func newTx(ctx *cli.Context) (err error) {
	cfg, err := common.GetNetworkConfig(ctx)
	if err != nil {
		return err
	}

	messageSlice := ctx.StringSlice("messages")

	msgs := make([]sdk.Msg, len(messageSlice))

	for idx, msgHex := range messageSlice {
		buf, err := hex.DecodeString(msgHex)
		if err != nil {
			return err
		}

		msg, err := messages.Unmarshal(cfg.EncodingConfig(), buf)
		if err != nil {
			return fmt.Errorf("invalid message: %v", err)
		}

		msgs[idx] = msg
	}

	gasLimit := ctx.Uint64("gas_limit")
	fee := sdk.NewCoins(sdk.NewInt64Coin("uumee", ctx.Int64("fee")))
	timeoutHeight := ctx.Uint64("timeout_height")
	memo := ctx.String("memo")

	pkHex, err := hex.DecodeString(ctx.String("acc_pubkey"))
	if err != nil {
		return err
	}

	pk, err := account.Secp256k1PublicKey(pkHex)
	if err != nil {
		return err
	}

	acc, err := account.NewUserAccount(cfg, pk, ctx.Uint64("acc_number"), ctx.Uint64("acc_sequence"))
	if err != nil {
		return err
	}

	txnBuilder, err := transaction.Build(cfg, msgs, gasLimit, fee, memo, timeoutHeight, []*account.Account{acc})
	if err != nil {
		return fmt.Errorf("could not build transaction : %v", err)
	}

	serializedTxn, err := transaction.Serialize(cfg.EncodingConfig().TxConfig, txnBuilder.GetTx())
	if err != nil {
		return fmt.Errorf("could not serialize transaction : %v", err)
	}

	fmt.Printf("%x\n", serializedTxn)
	return nil
}

func signTx(ctx *cli.Context) (err error) {
	cfg, err := common.GetNetworkConfig(ctx)
	if err != nil {
		return err
	}

	txHex, err := hex.DecodeString(ctx.String("tx"))
	if err != nil {
		return err
	}

	tx, err := transaction.FromBytes(cfg.EncodingConfig(), txHex)
	if err != nil {
		return fmt.Errorf("unable to deserialize : %v", err)
	}

	builder, err := transaction.WrapBuilder(cfg.EncodingConfig().TxConfig, tx)
	if err != nil {
		return fmt.Errorf("could not deserialize transaction : %v", err)
	}

	signerData := transaction.NewSignerData(ctx.String("chain_id"), ctx.Uint64("acc_number"), ctx.Uint64("acc_sequence"))

	sk, err := account.PrivateKeyFromHex(ctx.String("private_key"))
	if err != nil {
		return err
	}

	txn, err := transaction.Sign(cfg.EncodingConfig().TxConfig, builder, *signerData, sk)
	if err != nil {
		return err
	}

	serializedTxn, err := transaction.Serialize(cfg.EncodingConfig().TxConfig, txn)
	if err != nil {
		return fmt.Errorf("unable to serialize : %v", err)
	}

	fmt.Printf("%x\n", serializedTxn)
	return nil
}
