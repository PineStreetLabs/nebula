# Getting Started

## Running a local node
Nebula can connect and interact with a local node using RPC.

For example, to interact with the Umee network you will need to access an Umee node.

Running an Umee node is easy with [Starport](https://github.com/tendermint/starport).

1. Install Starport using `go install`

```sh
go install github.com/tendermint/starport@latest
```

2. Clone the Umee respoitory

```sh
git clone https://github.com/umee-network/umee.git && cd umee
```

2. Use Starport to serve the chain

```sh
starport chain serve -c starport.ci.yml
```

## Usage

```sh
  # create a local binary
  make build
  # alternatively, install to the GOBIN path
  go install -v ./...
```

## Send & Receive

Create a new account.

```sh
$ export ACCOUNT=$(./nebula --network=umee account)
$ echo $ACCOUNT | jq
{
  "address": "umee1ldxhrcu4vpr4xcgaffs587j5zunul8gu2c9cxd",
  "private_key": "9c9038a9bfd0bba17ac0eb709bb1db88c8f12663e0e93d11e5694b0fec0f5842"
}
```

Fund the account using a faucet. Assuming you are running `starport`, you might have a faucet available at localhost:4500.

```sh
$ export ADDRESS=$(echo $ACCOUNT | jq -r .address)
$ curl -X POST http://localhost:4500 -d '{"address":"'"$ADDRESS"'","coins":["2000uumee"]}'
```

Check the balance and retrieve the Account Number.

```sh
$ curl http://localhost:1317/auth/accounts/$ADDRESS
{
  "height": "25812",
  "result": {
    "type": "cosmos-sdk/BaseAccount",
    "value": {
      "address": "umee1ldxhrcu4vpr4xcgaffs587j5zunul8gu2c9cxd",
      "account_number": "13"
    }
  }
}
$ export ACC_NUMBER=$(curl http://localhost:1317/auth/accounts/$ADDRESS | jq -r .result.value.account_number)
13
```

Alternatively, check the balance and Account number using nebula.
```sh
$ nebula --network=umee balance -address=$ADDRESS
2000uumee
$ nebula --network=umee account_info -address=$ADDRESS | jq
{
  "address": "umee1x7kxvhzruz3q7tlt5x5hx2yjgqyzk37wkmlljv",
  "publickey": "",
  "sequence": 0,
  "number": 13,
}
```

Try sending a transaction.

```sh
$ export SK=$(echo $ACCOUNT | jq -r .private_key)
$ nebula --network=umee bank_send --recipient=$ADDRESS --amount=1000 --fee=1 --gas_limit=400000 --timeout_height=100000 --private_key=$SK --acc_number=$ACC_NUMBER --acc_sequence=0 --memo="bank send"
```

Umee Leverage transactions
```sh
$ export SK=$(echo $ACCOUNT | jq -r .private_key)
$ nebula --network=umee lend_asset --amount=100 --fee=1 --gas_limit=400000 --timeout_height=100000 --private_key=$SK --acc_number=$ACC_NUMBER --acc_sequence=0 --memo="lend asset"
$ nebula --network=umee withdraw_asset --amount=100 --fee=1 --gas_limit=400000 --timeout_height=1000000 --private_key=$SK --acc_number=$ACC_NUMBER --acc_sequence=0 --memo="withdraw asset"
$ nebula --network=umee set_collateral --enabled=true --fee=1 --gas_limit=400000 --timeout_height=1000000 --private_key=$SK --acc_number=$ACC_NUMBER --acc_sequence=0 --memo="set collateral"
$ nebula --network=umee repay_asset --amount=100 --fee=1 --gas_limit=400000 --timeout_height=1000000 --private_key=$SK --acc_number=$ACC_NUMBER --acc_sequence=0 --memo="repay asset"
```