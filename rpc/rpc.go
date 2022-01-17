package rpc

import (
	"context"
	"net/http"

	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	tendermintHttp "github.com/tendermint/tendermint/rpc/client/http"
)

// Config holds the configuration for the RPC client.
// Each node exposes endpoints to gRPC, REST, and Tendermint RPC servers.
type Config struct {
	grpcAddress string
	rpcAddress  string
}

// NewConfig returns a new Config.
func NewConfig(grpcAddress, rpcAddress string) *Config {
	return &Config{
		grpcAddress: grpcAddress,
		rpcAddress:  rpcAddress,
	}
}

// Client maintains a connection to a node daemon.
type Client struct {
	cfg        *Config
	rpcClient  *tendermintHttp.HTTP
	grpcClient *grpc.ClientConn
}

// NewClient creates a new RPC client.
func NewClient(cfg *Config) (*Client, error) {
	httpClient := &http.Client{}

	// todo, right path?
	rpcClient, err := tendermintHttp.NewWithClient(cfg.rpcAddress, "/", httpClient)
	if err != nil {
		return nil, err
	}

	grpcClient, err := grpc.Dial(cfg.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		rpcClient:  rpcClient,
		grpcClient: grpcClient,
		cfg:        cfg,
	}, nil
}

// BroadcastTransaction broadcasta a transaction using a node daemon.
func (c *Client) BroadcastTransaction(transaction []byte) (*coretypes.ResultBroadcastTx, error) {
	resp, err := c.rpcClient.BroadcastTxSync(context.Background(), transaction)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Balance returns the account balance for a given network.
func (c *Client) Balance(ncfg *networks.Params, address string) (*types.Coin, error) {
	client := bankTypes.NewQueryClient(c.grpcClient)

	resp, err := client.Balance(context.Background(), &bankTypes.QueryBalanceRequest{
		Address: address,
		Denom:   ncfg.Denom(),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetBalance(), nil
}
