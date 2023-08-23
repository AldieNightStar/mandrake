package mandrake

import (
	"testing"
)

func TestBase64(t *testing.T) {
	original := "Hello!"
	b64 := to64([]byte(original))
	current, _ := from64(b64)

	if string(current) != original {
		t.Fatal("(1) Bad 64 decode: ", string(current))
	}

	original = "This is the right spot"
	b64 = to64([]byte(original))
	current, _ = from64(b64)

	if string(current) != original {
		t.Fatal("(2) Bad 64 decode: ", string(current))
	}
}
