# Nebula
Nebula is unified wallet interface for Cosmos. This library is designed to support any app-chain within the Cosmos ecosystem by relying on module interfaces to handle message implementations. Signing messages is done offline by delivering a payload to your HSM or key management software. Alternatively, it is possible to use in-memory key management but it is highly unadvisable.

This software makes it easy for custodians, exchanges, and engineers to build wallet applications on top of the Cosmos network.

## Background
Each app chain in the Cosmos ecosystem is defined by a series of modules. Modules implement the query and message capabilitiy of the chain. Messages are composed into tranasctions and finalized by each new block to the chain.
