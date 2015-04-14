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
