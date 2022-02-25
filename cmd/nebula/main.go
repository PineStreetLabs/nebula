package main

import (
	"errors"
	"log"
	"os"

	"github.com/PineStreetLabs/nebula/cmd/nebula/umee"
	"github.com/PineStreetLabs/nebula/rpc"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nebula"
	app.Usage = "Gateway to the Cosmos."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpc",
			Usage: "the host:port endpoint of the Tendermint RPC server (e.g. 127.0.0.1:26657)",
		},
		cli.StringFlag{
			Name:  "grpc",
			Usage: "the host:port endpoint of the gRPC sever (e.g. 127.0.0.1:9090)",
		},
		cli.StringFlag{
			Name:  "network",
			Usage: "network parameters",
		},
	}
	app.Commands = []cli.Command{
		newAccountCommand,
		newBankSendCommand,
		newTxCommand,
		signTxCommand,
		broadcastTxCommand,
		balanceCommand,
		accountCommand,
		bestBlockHeightCommand,
		blockByHashCommand,
		blockByHeightCommand,
		transactionCommand,
		umee.LendAssetCommand,
		umee.WithdrawAssetCommand,
		umee.SetCollateralCommand,
		umee.RepayAssetCommand,
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

func getClient(p *profile) (*rpc.Client, error) {
	if p == nil {
		return nil, errors.New("server endpoints not supplied")
	}

	return rpc.NewClient(rpc.NewConfig(p.grpc, p.rpc))
}

func rpcClient(ctx *cli.Context) (*rpc.Client, error) {
	profile, err := getProfile(ctx)
	if err != nil {
		return nil, err
	}

	return getClient(profile)
}
