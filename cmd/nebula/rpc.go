package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
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

	fmt.Printf("%v\n", response)
	return nil
}

func queryBalance(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	params, err := getNetworkConfig(ctx)
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

	params, err := getNetworkConfig(ctx)
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

func simulateTx(grpcConn *grpc.ClientConn, txBytes []byte) error {
	txClient := tx.NewServiceClient(grpcConn)
	grpcRes, err := txClient.Simulate(
		context.Background(),
		&tx.SimulateRequest{
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(grpcRes.GasInfo)
	return nil
}
