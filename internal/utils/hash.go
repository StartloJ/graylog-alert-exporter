// Package utils provides utility function to help and reduce dumplicate code
package utils

import (
	"crypto/sha256"
	"fmt"
)

// Hash use sha256 to hash and return in string format
func Hash(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
