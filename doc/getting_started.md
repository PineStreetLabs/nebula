# Getting Started


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
$ nebula bank_send --network=umee --recipient=$ADDRESS --fee=1 --gas_limit=400000 --timeout_height=100000 --private_key=$SK --acc_number=$ACC_NUMBER --acc_sequence=0 --memo=""
```
