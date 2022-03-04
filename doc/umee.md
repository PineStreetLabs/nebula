# Umee

Nebula natively supports [Umee](https://www.umee.cc).

## Testnet
The Umee network has a range of testnets and endpoints available [here](https://github.com/umee-network/testnets/tree/main/networks). Support for each testnet might vary.

## Mainnet
A [list](https://github.com/umee-network/umee/tree/main/networks/umee-1) of mainnet Umee archival nodes. 

## Usage

```sh
$ export RPC=https://rpc.blue.main.network.umee.cc
$ export NETWORK=umee
$ nebula --rpc=$RPC --network=$NETWORK bestblockheight
116813
```
Request account information:
```sh
$ nebula --rpc=$RPC --network=$NETWORK account_info --address=umee1vp7yuww2lenznk6tjv80gr4258mzmk56gf9xtm
```
```json
{"address":"umee1vp7yuww2lenznk6tjv80gr4258mzmk56gf9xtm","publickey":"","sequence":0,"number":93245}
```
Create a new account:
```sh
$ nebula --network=$NETWORK account
```
```json
{"address":"umee14vcvlugaag99uac496acpdxuwapjysa2wrrcnd","private_key":"..."}
```

In order to craft a transaction that sends uumee to recipient we need to (1) create the Bank module's MsgSend message, (2) combine the message into a transaction, (3) sign the transaction, and (4) broadcast the transaction.

(1)
```sh
$ nebula --network=$NETWORK bank_send --recipient=umee14vcvlugaag99uac496acpdxuwapjysa2wrrcnd --amount=1 --sender=umee1vp7yuww2lenznk6tjv80gr4258mzmk56gf9xtm
```

(2)
```sh
$ nebula --network=$NETWORK new_tx --messages=$MSG --acc_pubkey=02452611abd6595aefec1889a0244c28ebeb78e1fa490e1d61f6af1f3d7722899d --fee=0 --gas_limit=80000 --timeout_height=0 --memo=""
```

(3)
```sh
$ nebula --network=$NETWORK sign_tx --tx=$TX --private_key=$SK --chain_id="umee-1" --acc_number=93245 --acc_sequence=0
```

(4)
```sh
$ nebula --rpc=$RPC --network=$NETWORK broadcast_tx --tx_hex=$TX_HEX
```

Check for the transaction status
```sh
$ nebula --rpc=$RPC --network=$NETWORK transaction --txid=1FF455ED2974F0E6B6A954166706F506704CA4610890E53C1FD3CA5E8945D614 | jq
```
```json
{
  "hash": "1FF455ED2974F0E6B6A954166706F506704CA4610890E53C1FD3CA5E8945D614",
  "height": 134218,
  "index": 7,
    ...
}
```