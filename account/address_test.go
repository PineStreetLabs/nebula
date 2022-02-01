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

		// Test marshalling
		{
			buf, err := addr.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			newAddr := Address{}
			if err := (&newAddr).Unmarshal(buf); err != nil {
				t.Fatal(err)
			}

			if !newAddr.Equals(addr) {
				t.Fatal("expected equal addresses")
			}
		}

		// Test JSON marshalling
		{
			buf, err := addr.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}

			newAddr := Address{}
			if err := (&newAddr).UnmarshalJSON(buf); err != nil {
				t.Fatal(err)
			}

			if !newAddr.Equals(addr) {
				t.Fatal("expected equal addresses")
			}
		}

		// Test YAML marshalling.
		{
			buf, err := addr.MarshalYAML()
			if err != nil {
				t.Fatal(err)
			}

			bufStr := (buf).(string)

			newAddr := Address{}
			if err := (&newAddr).UnmarshalYAML([]byte(bufStr)); err != nil {
				t.Fatal(err)
			}

			t.Log(newAddr)
			t.Log(addr)
			if !newAddr.Equals(addr) {
				t.Fatal("expected equal addresses")
			}
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
