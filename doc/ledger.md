# Ledger Support

The Nebula CLI supports signing from Ledger hardware wallet devices.

## Account
Create an account using a private key protected by your Ledger.

```sh
nebula --network $NETWORK account --ledger_account 0 --ledger
```
```json
{
  "address": "umee1tr7v5jtph2pq9ceu04y8ttx7j9v272npgs097j",
  "publickey": "033af5774296c0c566ec4866c20a1233d7f15cccc4580f1dd73901bc8d14d85984",
  "sequence": 0,
  "number": 0,
  "privatekey": "7901bfa2c35a252739edbb87eda344e5318fb40720c4737449bf43f788eee819"
}
```

## Sign
Partially sign a transaction using a private key found on your Ledger. This is particularly useful for signing transactions for multi-signature accounts.

```sh
nebula --network=cosmos partial_sign_tx --tx=$TX --ledger_account 0
```