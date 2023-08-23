package mandrake

import (
	"encoding/base64"
)

func to64(dat []byte) string {
	return base64.StdEncoding.EncodeToString(dat)
}

func from64(s string) ([]byte, error) {
	dat, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func EncodeBase64(dat []byte, password []byte) string {
	enc := Encode(dat, password)
	return to64(enc)
}

func DecodeBase64(str string, password []byte) ([]byte, error) {
	dat, err := from64(str)
	if err != nil {
		return nil, err
	}
	return Decode(dat, password), nil
}
