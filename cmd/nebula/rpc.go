package main

import (
	"context"
	"encoding/hex"
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

	response, err := rpcClient.BroadcastTransaction(txBytes)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", response)
	return nil
}

func balance(ctx *cli.Context) error {
	rpcClient, err := rpcClient(ctx)
	if err != nil {
		return err
	}

	params, err := getNetworkConfig(ctx)
	if err != nil {

	}

	response, err := rpcClient.Balance(params, ctx.String("address"))
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
