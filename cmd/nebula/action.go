package main

import (
	"errors"
	"fmt"
	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/urfave/cli"
)

var errUnsupportedNetwork = errors.New("unsupported network")

func newAccount(ctx *cli.Context) (err error) {
	sk := account.NewPrivateKey()

	var cfg *networks.Params
	switch network := ctx.String("network"); network {
	case networks.Cosmos:
		cfg = networks.GetCosmosCfg()
	case networks.Umee:
		cfg = networks.GetUmeeCfg()
	default:
		return fmt.Errorf("%w : %s", errUnsupportedNetwork, network)
	}

	acc, err := account.FromPublicKey(cfg, sk.PubKey())
	if err != nil {
		return err
	}

	fmt.Printf("%x\n", sk.Bytes())
	fmt.Printf("%s\n", acc.GetAddress().String())
	return nil
}
