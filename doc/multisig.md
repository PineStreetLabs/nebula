# MultiSig 

Nebula supports creating MultiSig accounts and signing transcations with MultiSig accounts as signers.

## Usage

Create two accounts.
```sh
nebula --network=cosmos account
...
export PK0=03befb7859171a25b6a338887de6aa6f92128aa5cde19a7504fdba69df79d58ac7
export PK1=02161503851ca8bc599d60fd51f0ea8b0c30d41f30e3e1378db7febedc16812539
```

Create a MultiSig account using the public keys from the two accounts.
```sh
nebula --network=cosmos multisig_account --threshold=1 --publickey=$PK0 --publickey=$PK1
```
```json
{
  "address": "cosmos12j23rlqllfpmjuq4zkepesntadksv60uv3tgf6",
  "publickey": "22c1f7e208011226eb5ae9872103befb7859171a25b6a338887de6aa6f92128aa5cde19a7504fdba69df79d58ac71226eb5ae9872102161503851ca8bc599d60fd51f0ea8b0c30d41f30e3e1378db7febedc16812539",
  "sequence": 0,
  "number": 0
}
{
  "threshold": 1,
  "public_keys": [
    {
      "@type": "/cosmos.crypto.secp256k1.PubKey",
      "key": "A777eFkXGiW2oziIfeaqb5ISiqXN4Zp1BP26ad951YrH"
    },
    {
      "@type": "/cosmos.crypto.secp256k1.PubKey",
      "key": "AhYVA4UcqLxZnWD9UfDqiwww1B8w4+E3jbf+vtwWgSU5"
    }
  ]
}
```

To sign a transaction with a MultiSig signer, the following steps must be taken:

(1) parially sign the transaction
```sh
nebula --network=cosmos partial_sign_tx --tx=$TX ...
```
(2) combine signatures 
```sh
nebula --network=cosmos combine_signatures --tx=$TX --multisig_account=$MULTISIG_JSON --signature=$SIG0 --signature=$SIG1
```