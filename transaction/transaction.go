package transaction

import (
	"errors"
	"fmt"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// Build crafts transaction for a given network.
func Build(cfg *networks.Params, msgs []sdk.Msg, gasLimit uint64, fees sdk.Coins, memo string, timeoutHeight uint64) (client.TxBuilder, error) {
	// Build the transaction.
	builder := cfg.EncodingConfig().TxConfig.NewTxBuilder()
	if err := builder.SetMsgs(msgs...); err != nil {
		return nil, fmt.Errorf("add messages : %v", err)
	}

	builder.SetGasLimit(gasLimit)
	builder.SetFeeAmount(fees)
	builder.SetMemo(memo)
	builder.SetTimeoutHeight(timeoutHeight)

	return builder, nil
}

// NewSignerData creates SignerData that is required for signature creation.
func NewSignerData(chainID string, accNumber, accSeq uint64) *authsigning.SignerData {
	return &authsigning.SignerData{
		ChainID:       chainID,
		AccountNumber: accNumber,
		Sequence:      accSeq,
	}
}

// WrapBuilder accepts an sdk.Tx and returns a TxBuilder.
// Useful for signing.
func WrapBuilder(cfg client.TxConfig, tx sdk.Tx) (client.TxBuilder, error) {
	return cfg.WrapTxBuilder(tx)
}

// CombineSignatures combines signatures and finalizes a transaction for a multisig account.
func CombineSignatures(cfg client.TxConfig, txn client.TxBuilder, signatures []signingtypes.SignatureV2, account *account.Account) (signing.Tx, error) {
	var multisigPub *kmultisig.LegacyAminoPubKey
	if pubkey, ok := account.GetPubKey().(*kmultisig.LegacyAminoPubKey); ok {
		multisigPub = pubkey
	} else {
		return nil, errors.New("expected multisig account")
	}

	multisigInfo := multisig.NewMultisig(len(multisigPub.PubKeys))

	for _, sig := range signatures {
		if err := multisig.AddSignatureV2(multisigInfo, sig, multisigPub.GetPubKeys()); err != nil {
			return nil, err
		}
	}

	sig := signingtypes.SignatureV2{
		PubKey:   multisigPub,
		Data:     multisigInfo,
		Sequence: account.GetSequence(),
	}

	if err := txn.SetSignatures(sig); err != nil {
		return nil, fmt.Errorf("set signatures : %v", err)
	}

	if err := txn.GetTx().ValidateBasic(); err != nil {
		return nil, err
	}

	return txn.GetTx(), nil
}

// PartialSign creates a signature for multisig accounts that have multiple signers.
func PartialSign(cfg client.TxConfig, txn client.TxBuilder, signerData authsigning.SignerData, sk cryptotypes.PrivKey) (signingtypes.SignatureData, error) {
	signMode := signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON

	signBytes, err := cfg.SignModeHandler().GetSignBytes(signMode, signerData, txn.GetTx())
	if err != nil {
		return nil, err
	}

	sig, err := sk.Sign(signBytes)
	if err != nil {
		return nil, err
	}

	signature := &signingtypes.SingleSignatureData{SignMode: signMode, Signature: sig}

	// Verify
	if err := signing.VerifySignature(sk.PubKey(), signerData, signature, cfg.SignModeHandler(), txn.GetTx()); err != nil {
		return nil, err
	}

	return signature, nil
}

// Sign accepts a valid transaction and signs it.
func Sign(cfg client.TxConfig, txn client.TxBuilder, signerData authsigning.SignerData, sk cryptotypes.PrivKey) (signing.Tx, error) {
	sig, err := tx.SignWithPrivKey(cfg.SignModeHandler().DefaultMode(), signerData, txn, sk, cfg, signerData.Sequence)
	if err != nil {
		return nil, fmt.Errorf("unable to sign : %v", err)
	}

	if err := txn.SetSignatures(sig); err != nil {
		return nil, fmt.Errorf("set signatures : %v", err)
	}

	return txn.GetTx(), nil
}

// FromBytes deserializes a transaction.
func FromBytes(cfg networks.EncodingConfig, txn []byte) (sdk.Tx, error) {
	return cfg.TxConfig.TxDecoder()(txn)
}

// FromJSON deserializes a transaction.
func FromJSON(cfg client.TxConfig, txn []byte) (sdk.Tx, error) {
	return cfg.TxJSONDecoder()(txn)
}

// Serialize serializes a transaction to bytes.
func Serialize(cfg client.TxConfig, txn signing.Tx) ([]byte, error) {
	return cfg.TxEncoder()(txn)
}

// SerializeJSON serializes a transaction to JSON bytes.
func SerializeJSON(cfg client.TxConfig, txn sdk.Tx) ([]byte, error) {
	return cfg.TxJSONEncoder()(txn)
}
