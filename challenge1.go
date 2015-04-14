package set1

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
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
	decrypted := xor(b, b2)
	return []byte(hex.EncodeToString(decrypted))
}

func xor(key []byte, encrypted []byte) []byte {
	prod := make([]byte, len(encrypted))
	for i := 0; i < len(prod); i++ {
		prod[i] = key[i%len(key)] ^ encrypted[i]
	}
	return prod
}

func probablyEnglish(s string) bool {
	letterFrequencies := map[string]float64{
		"e": 12.575645,
		"t": 9.085226,
		"a": 8.000395,
		"o": 7.591270,
		"i": 6.920007,
		"n": 6.903785,
		"s": 6.340880,
		"h": 6.236609,
		"r": 5.959034,
		"d": 4.317924,
		"l": 4.057231,
		"u": 2.841783,
		"c": 2.575785,
		"m": 2.560994,
		"f": 2.350463,
		"w": 2.224893,
		"g": 1.982677,
		"y": 1.900888,
		"p": 1.795742,
		"b": 1.535701,
		"v": 0.981717,
		"k": 0.739906,
		"x": 0.179556,
		"j": 0.145188,
		"q": 0.117571,
		"z": 0.079130,
	}
	allowedEnglishCharacters := `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ., '":;!$£%`
	totalScore := 0.00
	strings.Replace(s, " ", "", -1)
	for _, r := range s {
		if strings.Index(allowedEnglishCharacters, string(r)) == -1 {
			return false
		}
		totalScore = totalScore + letterFrequencies[string(r)]
	}
	arbitraryScoreCutoff := float64(len(s) * 2)
	return totalScore > arbitraryScoreCutoff
}

func decryptHex(encryptedHex []byte) string {
	b, err := hex.DecodeString(string(encryptedHex))
	if err != nil {
		panic(err)
	}
	possibleKeys := []byte("ABCDEFGHIJKLMNOPQRSSTUVWXYZ")
	for i := 0; i < len(possibleKeys); i++ {
		decryptedHex := string(xor([]byte(string(possibleKeys[i])), b))
		if probablyEnglish(decryptedHex) {
			return decryptedHex
		}
	}
	panic("no match found")
}
