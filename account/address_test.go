package account

import (
	"bytes"
	"testing"
)

func TestAddressFromString(t *testing.T) {
	{
		addr, err := AddressFromString("prefixa1qqqsyqcyq5rqwzqfpg9scrgwpugpzysn7hzdtn")
		if err != nil {
			t.Fatal(err)
		}

		expectedHrp := "prefixa"
		if addr.hrp != expectedHrp {
			t.Fatalf("hrp %s != %s", addr.hrp, expectedHrp)
		}

		expectedBuf := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
		if !bytes.Equal(addr.data, expectedBuf) {
			t.Fatalf("%v != %v", addr.data, expectedBuf)
		}
	}

	{
		addr, err := AddressFromString("prefixb1qqqsyqcyq5rqwzqf20xxpc")
		if err != nil {
			t.Fatal(err)
		}

		expectedHrp := "prefixb"
		if addr.hrp != expectedHrp {
			t.Fatalf("hrp %s != %s", addr.hrp, expectedHrp)
		}

		expectedBuf := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !bytes.Equal(addr.data, expectedBuf) {
			t.Fatalf("%v != %v", addr.data, expectedBuf)
		}
	}

}
