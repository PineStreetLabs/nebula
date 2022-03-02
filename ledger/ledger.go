package ledger

import (
	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/btcsuite/btcd/btcec"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tendermint_btcec "github.com/tendermint/btcd/btcec"

	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	ledger_cosmos_go "github.com/cosmos/ledger-cosmos-go"
)

// Account returns an Account from a Ledger device.
func Account(cfg *networks.Params, accountBIP44 uint32) (*account.Account, error) {
	device, err := ledger_cosmos_go.FindLedgerCosmosUserApp()
	if err != nil {
		return nil, err
	}

	path := *hd.NewFundraiserParams(accountBIP44, sdk.CoinType, 0)
	pkBuf, addr, err := device.GetAddressPubKeySECP256K1(path.DerivationPath(), cfg.AccountHRP())
	if err != nil {
		return nil, err
	}

	pk, err := account.Secp256k1PublicKey(pkBuf)
	if err != nil {
		return nil, err
	}

	return account.NewAccount(addr, pk, 0, 0)
}

// Sign signs a transaction from a Ledger device. Returns a signature.
func Sign(cfg *networks.Params, accountBIP44 uint32, txn client.TxBuilder, signerData authsigning.SignerData) (*signingtypes.SignatureV2, error) {
	device, err := ledger_cosmos_go.FindLedgerCosmosUserApp()
	if err != nil {
		return nil, err
	}

	path := *hd.NewFundraiserParams(accountBIP44, sdk.CoinType, 0)

	pkBuf, err := device.GetPublicKeySECP256K1(path.DerivationPath())
	if err != nil {
		return nil, err
	}

	pk, err := account.Secp256k1PublicKey(pkBuf)
	if err != nil {
		return nil, err
	}

	signBytes, err := cfg.EncodingConfig().TxConfig.SignModeHandler().GetSignBytes(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signerData, txn.GetTx())
	if err != nil {
		return nil, err
	}

	signatureDER, err := device.SignSECP256K1(path.DerivationPath(), signBytes)
	if err != nil {
		return nil, err
	}

	sigDER, err := btcec.ParseDERSignature(signatureDER, btcec.S256())
	if err != nil {
		return nil, err
	}

	sigBER := tendermint_btcec.Signature{R: sigDER.R, S: sigDER.S}

	sigData := &signingtypes.SingleSignatureData{
		SignMode:  signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
		Signature: sigBER.Serialize(),
	}

	if err := authsigning.VerifySignature(pk, signerData, sigData, cfg.EncodingConfig().TxConfig.SignModeHandler(), txn.GetTx()); err != nil {
		return nil, err
	}

	return &signingtypes.SignatureV2{PubKey: pk, Data: sigData}, nil
}
