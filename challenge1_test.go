package set1

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestConvertsHexToBase64(t *testing.T) {
	b := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	assert.Equal(t, []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"), convertToBase64(b))
}
