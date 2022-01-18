package rpc

import (
	"context"
	"net/http"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
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

// BroadcastTransaction broadcasts a transaction using a node daemon.
func (c *Client) BroadcastTransaction(ctx context.Context, transaction []byte) (*coretypes.ResultBroadcastTx, error) {
	resp, err := c.rpcClient.BroadcastTxSync(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Balance returns the account balance for a given network.
func (c *Client) Balance(ctx context.Context, ncfg *networks.Params, address string) (*types.Coin, error) {
	client := bankTypes.NewQueryClient(c.grpcClient)

	resp, err := client.Balance(ctx, &bankTypes.QueryBalanceRequest{
		Address: address,
		Denom:   ncfg.Denom(),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetBalance(), nil
}

// Transaction returns the transaction using its hash.
func (c *Client) Transaction(ctx context.Context, txID []byte) (*coretypes.ResultTx, error) {
	return c.rpcClient.Tx(ctx, txID, true)
}

// BestBlockHeight returns the current network block height.
func (c *Client) BestBlockHeight(ctx context.Context) (int64, error) {
	status, err := c.rpcClient.Status(ctx)
	if err != nil {
		return 0, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}

// BlockByHeight returns a block given its height.
func (c *Client) BlockByHeight(ctx context.Context, height int64) (*coretypes.ResultBlock, error) {
	return c.rpcClient.Block(ctx, &height)
}

// BlockByHash returns a block given its hash.
func (c *Client) BlockByHash(ctx context.Context, hash []byte) (*coretypes.ResultBlock, error) {
	return c.rpcClient.BlockByHash(ctx, hash)
}

// Account returns account details.
func (c *Client) Account(ctx context.Context, nCfg *networks.Params, address string) (*account.Account, error) {
	client := authtypes.NewQueryClient(c.grpcClient)
	resp, err := client.Account(ctx, &authtypes.QueryAccountRequest{
		Address: address,
	})
	if err != nil {
		return nil, err
	}

	var acc authtypes.AccountI
	if err := nCfg.EncodingConfig().InterfaceRegistry.UnpackAny(resp.Account, &acc); err != nil {
		return nil, err
	}

	return account.NewAccount(acc.GetAddress().String(), acc.GetPubKey(), acc.GetAccountNumber(), acc.GetSequence())
}
