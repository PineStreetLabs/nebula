package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	protoEnc "google.golang.org/grpc/encoding/proto"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tendermintHttp "github.com/tendermint/tendermint/rpc/client/http"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
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
	rpcClient  *rpcClient
	grpcClient *grpc.ClientConn
}

type rpcClient struct {
	*tendermintHttp.HTTP
}

var errEndpointsUnavailable = errors.New("client connection unavailable")
var errGrpcUnavailable = errors.New("grpc client connection unavailable")
var errRPCUnavailable = errors.New("rpc client connection unavailable")

func (c Client) getGrpcClient() (*grpc.ClientConn, error) {
	if c.grpcClient == nil {
		return nil, errGrpcUnavailable
	}
	return c.grpcClient, nil
}

func (c Client) getRPCClient() (*rpcClient, error) {
	if c.rpcClient == nil {
		return nil, errRPCUnavailable
	}

	return c.rpcClient, nil
}

// NewClient creates a new RPC client.
func NewClient(cfg *Config) (c *Client, err error) {
	httpClient := &http.Client{}

	c = &Client{}

	// Assign a rpcClient if the address is available.
	if len(cfg.rpcAddress) != 0 {
		rpc, err := tendermintHttp.NewWithClient(cfg.rpcAddress, "/", httpClient)
		if err != nil {
			return nil, err
		}

		c.rpcClient = &rpcClient{rpc}
	}

	// Assign a grpcClient if the address is available.
	if len(cfg.grpcAddress) != 0 {
		// We use grpc.WithInsecure() because Cosmos-SDK does not support any transport security method.
		c.grpcClient, err = grpc.Dial(cfg.grpcAddress, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// BroadcastTransaction broadcasts a transaction using a node daemon.
func (c *Client) BroadcastTransaction(ctx context.Context, transaction []byte) (*coretypes.ResultBroadcastTx, error) {
	client, err := c.getRPCClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.BroadcastTxSync(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Balance returns the account balance for a given network.
func (c *Client) Balance(ctx context.Context, ncfg *networks.Params, address string) (*types.Coin, error) {
	var client bankTypes.QueryClient

	// If we have a gRPC client available, choose it.
	// Otherwise, use the rpc address for gRPC with REST.
	if grpcClient, err := c.getGrpcClient(); err == nil {
		client = bankTypes.NewQueryClient(grpcClient)
	} else if rpcClient, err := c.getRPCClient(); err == nil {
		client = bankTypes.NewQueryClient(rpcClient)
	} else {
		return nil, errEndpointsUnavailable
	}

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
	client, err := c.getRPCClient()
	if err != nil {
		return 0, err
	}

	status, err := client.Status(ctx)
	if err != nil {
		return 0, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}

// BlockByHeight returns a block given its height.
func (c *Client) BlockByHeight(ctx context.Context, height int64) (*coretypes.ResultBlock, error) {
	client, err := c.getRPCClient()
	if err != nil {
		return nil, err
	}

	return client.Block(ctx, &height)
}

// BlockByHash returns a block given its hash.
func (c *Client) BlockByHash(ctx context.Context, hash []byte) (*coretypes.ResultBlock, error) {
	client, err := c.getRPCClient()
	if err != nil {
		return nil, err
	}

	return client.BlockByHash(ctx, hash)
}

// Account returns account details.
func (c *Client) Account(ctx context.Context, nCfg *networks.Params, address string) (*account.Account, error) {
	var client authtypes.QueryClient

	// If we have a gRPC client available, choose it.
	// Otherwise, use the rpc address for gRPC with REST.
	if grpcClient, err := c.getGrpcClient(); err == nil {
		client = authtypes.NewQueryClient(grpcClient)
	} else if rpcClient, err := c.getRPCClient(); err == nil {
		client = authtypes.NewQueryClient(rpcClient)
	} else {
		return nil, errEndpointsUnavailable
	}

	req := &authtypes.QueryAccountRequest{
		Address: address,
	}

	resp, err := client.Account(ctx, req)
	if err != nil {
		return nil, err
	}

	return queryAccounttoAccount(nCfg.EncodingConfig(), resp)
}

var errParseQueryAccountReponse = errors.New("cannot parse QueryAccountResponse")
var errTypeConv = errors.New("issue converting type")

func queryAccounttoAccount(encCfg networks.EncodingConfig, resp *authtypes.QueryAccountResponse) (*account.Account, error) {
	// Convert to JSON
	buf, err := encCfg.Marshaler.MarshalJSON(resp.Account)
	if err != nil {
		return nil, err
	}

	var v map[string]interface{}
	json.Unmarshal(buf, &v)

	var addr string
	var accNum uint64
	var accSeq uint64

	{
		iface, exists := v["address"]
		if !exists {
			return nil, errParseQueryAccountReponse
		}

		val, ok := iface.(string)
		if !ok {
			return nil, errTypeConv
		}

		addr = val
	}

	{
		iface, exists := v["account_number"]
		if !exists {
			return nil, errParseQueryAccountReponse
		}

		val, ok := iface.(string)
		if !ok {
			return nil, errTypeConv
		}

		accNum, err = strconv.ParseUint(val, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	{
		iface, exists := v["sequence"]
		if !exists {
			return nil, errParseQueryAccountReponse
		}

		val, ok := iface.(string)
		if !ok {
			return nil, errTypeConv
		}

		accSeq, err = strconv.ParseUint(val, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return account.NewAccount(addr, nil, accNum, accSeq)
}

// NewStream fulfills the grpc.Client interface.
func (rpc *rpcClient) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("stream unavailable")
}

// Invoke fulfills the grpc.Client interface using gRPC with REST.
func (rpc *rpcClient) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) (err error) {
	reqBuf, err := encoding.GetCodec(protoEnc.Name).Marshal(args)
	if err != nil {
		return err
	}

	resp, err := rpc.ABCIQueryWithOptions(ctx, method, reqBuf, rpcclient.ABCIQueryOptions{})
	if err != nil {
		return err
	}

	if !resp.Response.IsOK() {
		return fmt.Errorf("acbi response : %v", resp.Response.GetCode())
	}

	if err = encoding.GetCodec(protoEnc.Name).Unmarshal(resp.Response.Value, reply); err != nil {
		return err
	}

	return nil
}
