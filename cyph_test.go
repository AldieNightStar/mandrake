package mandrake

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	encoded := Encode(
		[]byte("Hello"),
		[]byte("pwd"),
	)

	decoded := Decode(encoded, []byte("pwd"))

	if string(decoded) != "Hello" {
		t.Fatal("Decoded is not original")
	}
}

func TestEncodeDecodeFail(t *testing.T) {
	encoded := Encode(
		[]byte("Hello"),
		[]byte("pwd"),
	)

	decoded := Decode(encoded, []byte("wrong"))

	if string(decoded) == "Hello" {
		t.Fatal("Decoded is compromised")
	}
}
