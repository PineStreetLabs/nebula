# Architecture
Nebula is a wallet SDK that empowers software engineers to build wallets for Cosmos-SDK based app chains. This library provides a unified API for wallet functionality. We avoid calling a node daemon for transaction construction, but provide functionality to query nodes using our CLI. All transaction construction and validation can be done offline while only relying on the network for broadcasting.

The entire lifecycle of a wallet can be managed using this library.

Nebula is divided into tightly scoped packages that allow the user to easily import the code into their existing projects.

Unfortunately, the dependency creep is very large in Cosmos-SDK based projects and we currently do not have a proposed solution for avoiding dependencies.

    /account -> account manipulation
    /keychain -> key derivation functionality
    /rpc -> query client
    /messages -> message definitions
    /networks -> network parameters
    /transaction -> transaction construction

# Usage
This library is tailored for wallet engineers that want to extend the functionality of their application to support Cosmos-based app chains. This includes managing node operations for staking, creating transactions, and managing accounts. 

## Remote signing
Nebula provides the scaffolding to manage remote signing workflows. This includes creating messages, combining them into a transaction, and then signing remotely before broadcasting.

One can do this programmatically using `transaction` and `messages` packages. Otherwise, one can use the Nebula CLI.

    # Create messagse
    nebula bank_send <...>

    # Combine messages into a transaction
    nebula new_tx <...> 

    # Either sign offline or provide a secret key to the CLI
    nebula sign_tx

    # Broadcast using a local or remote RPC connection
    nebula --rpc=$rpc broadcast_tx --tx_hex=<...>

## Wallet Operations
Nebula provides a framework for improving wallet operations for Cosmos-based app chain users.

The library and CLI can be used for improving staking operations, adding send/receive support to wallets, and interacting with particular message types like governance.

### Staking
Nebula supports creating staking and delegation message types that can be included in a transaction. Staking operators will find Nebula useful in their workflow to restake assets and creating delegation message types.

### Wallets
Exchanges, custodians, and wallets can use Nebula for adding send and receive support. This includes address generation as well as transaction construction.

### Governance
Nebula can help craft governance message types to broadcast governance proposals and signal votes.