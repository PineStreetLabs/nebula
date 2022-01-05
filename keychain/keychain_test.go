package keychain

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestFromSeed(t *testing.T) {
	// test vectors taken from https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vectors
	var tests = []struct {
		seed      string // hex seed
		chaincode string // hex chaincode
		key       string // hex secret key
		path      string
	}{
		{
			seed:      "000102030405060708090a0b0c0d0e0f",
			chaincode: "873dff81c02f525623fd1fe5167eac3a55a049de3d314bb42ee227ffed37d508",
			key:       "e8f32e723decf4051aefac8e2c93c9c5b214313817cdb01a1494b917c8436b35",
		},
	}

	for tc, test := range tests {
		seed, _ := hex.DecodeString(test.seed)

		// Create master BIP32 node from seed.
		key, err := FromSeed(seed)
		if err != nil {
			t.Fatalf("%d: %s", tc, err)
		}

		// Assert secret key matches.
		{
			sk, err := hex.DecodeString(test.key)
			if err != nil {
				t.Fatalf("%d: %s", tc, err)
			}

			if !bytes.Equal(key.key[:], sk) {
				t.Fatalf("%d : keys not equal", tc)
			}
		}

		// Assert chain code matches.
		{
			chaincode, err := hex.DecodeString(test.chaincode)
			if err != nil {
				t.Fatalf("%d: %s", tc, err)
			}

			if !bytes.Equal(key.chaincode[:], chaincode) {
				t.Fatalf("%d : chaincodes not equal", tc)
			}
		}
	}
}
