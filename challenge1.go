package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func convertToBase64(input []byte) []byte {
	b, err := hex.DecodeString(string(input))
	if err != nil {
		panic(err)
	}
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(encoded, b)
	return encoded
}

func fixedXOR(key []byte, secretString []byte) []byte {
	b, err := hex.DecodeString(string(key))
	if err != nil {
		panic(err)
	}
	b2, err := hex.DecodeString(string(secretString))
	if err != nil {
		panic(err)
	}
	prod := make([]byte, len(b2))
	for i := 0; i < len(prod); i++ {
		prod[i] = b[i] ^ b2[i]
	}
	return []byte(hex.EncodeToString(prod))
}
