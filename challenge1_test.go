package set1

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestConvertsHexToBase64(t *testing.T) {
	b := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	assert.Equal(t, []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"), convertToBase64(b))
}

func TestFixedXor(t *testing.T) {
	key := []byte("1c0111001f010100061a024b53535009181c")
	secret := []byte("686974207468652062756c6c277320657965")
	assert.Equal(t, []byte("746865206b696420646f6e277420706c6179"), fixedXOR(key, secret))
}

func TestSingleByteXOR(t *testing.T) {
	encrpytedHex := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	assert.Equal(t, "Cooking MC's like a pound of bacon", decryptSingleXORdHex(encrpytedHex))
}
