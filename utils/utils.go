package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword returns a SHA256 hash of the password.
// In production, use a more secure algorithm such as bcrypt.
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
