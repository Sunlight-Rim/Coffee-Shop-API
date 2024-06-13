package tools

import "crypto/sha256"

// SHA256 generates SHA-256 hash string.
func SHA256(text string) string {
	h := sha256.Sum256([]byte(text))
	return string(h[:])
}
