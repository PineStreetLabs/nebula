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
