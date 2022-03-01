# Nebula
Nebula is unified wallet interface for Cosmos. This library is designed to support any app-chain within the Cosmos ecosystem by relying on module interfaces to handle message implementations. Signing messages is done offline by delivering a payload to your HSM or key management software. Alternatively, it is possible to use in-memory key management but it is highly unadvisable.

This software makes it easy for custodians, exchanges, and engineers to build wallet applications on top of the Cosmos network.

## Background
Each app chain in the Cosmos ecosystem is defined by a series of modules. Modules implement the query and message capabilitiy of the chain. Messages are composed into tranasctions and finalized by each new block to the chain.


## Network Support
| Network      | Repository | Version | Documentation |
|-----------|-----------|-----------|-----------|
| Cosmos Hub (ATOM) | [Cosmos](https://github.com/cosmos/gaia)| - | - |
| Umee |[Umee](https://github.com/umee-network/umee) | v0.3.0| [Umee](/doc/umee.md)

## Usage
    $ go install -v ./...
    $ nebula help

Nebula is both a library for managing wallet workflows as well as a CLI.

Documentation for usage is available in `/doc`:

* [Getting Started](/doc/getting_started.md)
* [Architecture](doc/arch.md)
* [MultiSig Usage](doc/multisig.md)
* [Umee Guide](doc/umee.md)

## CLI
    NAME:
    nebula - Gateway to the Cosmos.

    USAGE:
    nebula [global options] command [command options] [arguments...]

    COMMANDS:
    help, h  Shows a list of commands or help for one command

    data:
        balance          <address>
        account_info     <address>
        bestblockheight  
        blockbyhash      <hash>
        blockbyheight    <height>
        transaction      <txid>

    umee:
        lend_asset      Create a lend asset transaction.
        withdraw_asset  Create a withdraw asset transaction.
        set_collateral  Create a set collateral transaction.
        repay_asset     Create a repay asset transaction.

    wallet:
        account       Create a new account.
        bank_send     Create a Bank module MsgSend message.
        new_tx        Combines a slice of messages into a new transaction.
        sign_tx       Sign a serialized transaction.
        broadcast_tx  Broadcast a transaction

    GLOBAL OPTIONS:
    --rpc value      the host:port endpoint of the Tendermint RPC server (e.g. 127.0.0.1:26657)
    --grpc value     the host:port endpoint of the gRPC sever (e.g. 127.0.0.1:9090)
    --network value  network parameters
    --help, -h       show help


## Testing
[Starport](https://github.com/tendermint/starport) is the dedicated tool for integration and e2e testing.

Starport allows local development with specific accounts funded at gensis. Make sure are running umee v0.3.0

    starport chain serve -c ./config.yml

Below are links to supported networks and their Starport configurations:
* [Cosmos Hub](https://github.com/cosmos/gaia/blob/main/config.yml)
* [Umee](https://github.com/umee-network/umee/blob/main/starport.ci.yml)