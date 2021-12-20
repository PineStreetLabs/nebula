package transaction

import (
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

func Build(cfg *networks.Params, msgs []sdk.Msg, gasLimit uint64, fees sdk.Coins, memo string, timeoutHeight uint64) (client.TxBuilder, error) {
	builder := cfg.EncodingConfig().TxConfig.NewTxBuilder()
	if err := builder.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	builder.SetGasLimit(gasLimit)
	builder.SetFeeAmount(fees)
	builder.SetMemo(memo)
	builder.SetTimeoutHeight(timeoutHeight)

	return builder, nil
}

func NewSignerData(chainID string, accNumber, seq uint64) *signing.SignerData {
	return &signing.SignerData{
		ChainID:       chainID,
		AccountNumber: accNumber,
		Sequence:      seq,
	}
}

func Sign(cfg client.TxConfig, txn client.TxBuilder, signerData signing.SignerData, sk cryptotypes.PrivKey) (signing.Tx, error) {
	sig, err := tx.SignWithPrivKey(cfg.SignModeHandler().DefaultMode(), signerData, txn, sk, cfg, 0)
	if err != nil {
		return nil, err
	}

	if err := txn.SetSignatures(sig); err != nil {
		return nil, err
	}

	return txn.GetTx(), nil
}

func FromBytes(cfg client.TxConfig, txn []byte) (sdk.Tx, error) {
	return cfg.TxDecoder()(txn)
}

func FromJSON(cfg client.TxConfig, txn []byte) (sdk.Tx, error) {
	return cfg.TxJSONDecoder()(txn)
}

func Serialize(cfg client.TxConfig, txn signing.Tx) ([]byte, error) {
	return cfg.TxEncoder()(txn)
}

func SerializeJSON(cfg client.TxConfig, txn signing.Tx) ([]byte, error) {
	return cfg.TxJSONEncoder()(txn)
}
