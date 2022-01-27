package transaction

import (
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// Build crafts transaction for a given network.
func Build(cfg *networks.Params, msgs []sdk.Msg, gasLimit uint64, fees sdk.Coins, memo string, timeoutHeight uint64, signerPubKeys []cryptotypes.PubKey) (client.TxBuilder, error) {
	builder := cfg.EncodingConfig().TxConfig.NewTxBuilder()
	if err := builder.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	builder.SetGasLimit(gasLimit)
	builder.SetFeeAmount(fees)
	builder.SetMemo(memo)
	builder.SetTimeoutHeight(timeoutHeight)

	// Assign signers.
	signers := make([]signingtypes.SignatureV2, len(signerPubKeys))

	for idx, pk := range signerPubKeys {
		signers[idx] = signingtypes.SignatureV2{
			PubKey: pk,
			Data: &signingtypes.SingleSignatureData{
				SignMode:  cfg.EncodingConfig().TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			// Sequence: 0,
		}
	}

	if err := builder.SetSignatures(signers...); err != nil {
		return nil, err
	}

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

// Sign accepts a valid transaction and signs it.
func Sign(cfg client.TxConfig, txn client.TxBuilder, signerData authsigning.SignerData, sk cryptotypes.PrivKey) (signing.Tx, error) {
	sig, err := tx.SignWithPrivKey(cfg.SignModeHandler().DefaultMode(), signerData, txn, sk, cfg, signerData.Sequence)
	if err != nil {
		return nil, err
	}

	if err := txn.SetSignatures(sig); err != nil {
		return nil, err
	}

	return txn.GetTx(), nil
}

// FromBytes deserializes a transaction.
func FromBytes(cfg client.TxConfig, txn []byte) (sdk.Tx, error) {
	return cfg.TxDecoder()(txn)
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
func SerializeJSON(cfg client.TxConfig, txn signing.Tx) ([]byte, error) {
	return cfg.TxJSONEncoder()(txn)
}
