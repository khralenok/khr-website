package utilities

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// This function generate a random token
func NewRawToken(n int) (string, error) {
	b := make([]byte, n)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

// This function generate SHA-256 token hash
func TokenHash(raw string) []byte {
	h := sha256.Sum256([]byte(raw))
	return h[:]
}
