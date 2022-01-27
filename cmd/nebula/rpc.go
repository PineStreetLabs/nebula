package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/PineStreetLabs/nebula/cmd/nebula/common"

	"github.com/urfave/cli"
)

func broadcastTx(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	txBytes, err := hex.DecodeString(ctx.String("tx_hex"))
	if err != nil {
		return err
	}

	response, err := rpcClient.BroadcastTransaction(context.Background(), txBytes)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(response)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

	return nil
}

func queryBalance(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	params, err := common.GetNetworkConfig(ctx)
	if err != nil {
		return err
	}

	response, err := rpcClient.Balance(context.Background(), params, ctx.String("address"))
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", response)

	return nil
}

func queryAccount(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	params, err := common.GetNetworkConfig(ctx)
	if err != nil {
		return err
	}

	response, err := rpcClient.Account(context.Background(), params, ctx.String("address"))
	if err != nil {
		return err
	}

	resp, err := json.Marshal(response)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

	return nil
}

func queryBestBlockheight(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	response, err := rpcClient.BestBlockHeight(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", response)

	return nil
}

func queryBlockByHeight(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	response, err := rpcClient.BlockByHeight(context.Background(), ctx.Int64("height"))
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", response)

	return nil
}

func queryBlockByHash(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	hash, err := hex.DecodeString(ctx.String("hash"))
	if err != nil {
		return err
	}

	response, err := rpcClient.BlockByHash(context.Background(), hash)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", response)

	return nil
}

func queryTransaction(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	hash, err := hex.DecodeString(ctx.String("txid"))
	if err != nil {
		return err
	}

	response, err := rpcClient.Transaction(context.Background(), hash)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(response)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp)

	return nil
}
