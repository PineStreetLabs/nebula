package ledger

import (
	"encoding/json"
	"fmt"

	"github.com/PineStreetLabs/nebula/account"
	"github.com/PineStreetLabs/nebula/networks"
	"github.com/PineStreetLabs/nebula/transaction"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
func Sign(cfg *networks.Params, accountBIP44 uint32, txn client.TxBuilder, signerData authsigning.SignerData) ([]byte, error) {
	device, err := ledger_cosmos_go.FindLedgerCosmosUserApp()
	if err != nil {
		return nil, err
	}

	txBuf, err := transaction.SerializeJSON(cfg.EncodingConfig().TxConfig, txn.GetTx())
	if err != nil {
		return nil, err
	}

	var txJSON map[string]map[string]json.RawMessage
	json.Unmarshal(txBuf, &txJSON)

	path := *hd.NewFundraiserParams(accountBIP44, sdk.CoinType, 0)

	// Ledger's Cosmos App requires a specific payload described here: https://github.com/cosmos/ledger-cosmos/blob/main/docs/TXSPEC.md
	payload := fmt.Sprintf(`{
		"account_number": %d,
		"chain_id": "%s",
		"fee": %s,
		"memo": "%s",
		"msgs": %s,
		"sequence": %d
	}`, signerData.AccountNumber, signerData.ChainID, txJSON["auth_info"]["fee"], txn.GetTx().GetMemo(), txJSON["body"]["messages"], signerData.Sequence)

	var val map[string]interface{}
	json.Unmarshal([]byte(payload), &val)
	buf, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	// Signing mode is SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	return device.SignSECP256K1(path.DerivationPath(), buf)
}
