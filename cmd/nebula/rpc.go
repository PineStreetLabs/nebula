package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"strconv"
)

func broadcastTx(ctx *cli.Context) error {
	host := ctx.String("host")
	port := ctx.Int("port")
	txBytes, err := hex.DecodeString(ctx.String("tx_hex"))
	if err != nil {
		return err
	}

	grpcConn, err := grpc.Dial(
		host+":"+strconv.Itoa(port),
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

	if err != nil {
		return err
	}

	return broadcast(grpcConn, txBytes)
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

func broadcast(grpcConn *grpc.ClientConn, txBytes []byte) error {
	txClient := tx.NewServiceClient(grpcConn)
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(grpcRes.TxResponse.Code) // Should be `0` if the tx is successful
	return nil
}
