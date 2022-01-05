# Nebula
Nebula is unified wallet interface for Cosmos. This library is designed to support any app-chain within the Cosmos ecosystem by relying on module interfaces to handle message implementations. Signing messages is done offline by delivering a payload to your HSM or key management software. Alternatively, it is possible to use in-memory key management but it is highly unadvisable.

This software makes it easy for custodians, exchanges, and engineers to build wallet applications on top of the Cosmos network.

## Background
Each app chain in the Cosmos ecosystem is defined by a series of modules. Modules implement the query and message capabilitiy of the chain. Messages are composed into tranasctions and finalized by each new block to the chain.


## Network Support
| Network      | Repository | Version |
| ----------- | ----------- | -----------|
| Cosmos Hub (ATOM) | [Cosmos](https://github.com/cosmos/gaia)| - |
| Umee |[Umee](https://github.com/umee-network/umee) | v0.3.0|

## Testing
[Starport](https://github.com/tendermint/starport) is the dedicated tool for integration and e2e testing.

Starport allows local development with specific accounts funded at gensis.

    starport chain serve -c ./config.yml

Below are links to supported networks and their Starport configurations:
* [Cosmos Hub](https://github.com/cosmos/gaia/blob/main/config.yml)
* [Umee](https://github.com/umee-network/umee/blob/main/starport.ci.yml)