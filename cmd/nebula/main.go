package main

import (
	"log"
	"os"

	rpc "github.com/PineStreetLabs/nebula/rpc"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nebula"
	app.Usage = "Gateway to the Cosmos."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpc",
			Value: "http://127.0.0.1:26657",
			Usage: "the host:port of the JSON-RPC server",
		},
		cli.StringFlag{
			Name:  "grpc",
			Value: "127.0.0.1:9090",
			Usage: "the host:port of the gRPC sever",
		},
		cli.StringFlag{
			Name:  "network",
			Usage: "network parameters",
		},
	}
	app.Commands = []cli.Command{
		broadcastTxCommand,
		newAccountCommand,
		newBankSendCommand,
		balanceCommand,
		accountCommand,
		bestBlockHeightCommand,
		blockByHashCommand,
		blockByHeightCommand,
		transactionCommand,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type profile struct {
	rpc  string
	grpc string
}

func getProfile(ctx *cli.Context) (*profile, error) {
	rpcAddress := ctx.GlobalString("rpc")
	grpcAddress := ctx.GlobalString("grpc")
	return &profile{
		rpc:  rpcAddress,
		grpc: grpcAddress,
	}, nil
}

func getClient(rpcAddress, grpcAddress string) (*rpc.Client, error) {
	return rpc.NewClient(rpc.NewConfig(grpcAddress, rpcAddress))
}

func rpcClient(ctx *cli.Context) (*rpc.Client, error) {
	profile, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}

	return getClient(profile.rpc, profile.grpc)
}
