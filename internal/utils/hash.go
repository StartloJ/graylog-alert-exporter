package utils

import "crypto/sha256"

// Hash use sha256 to hash and return in string format
func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return string(h.Sum(nil))
}
