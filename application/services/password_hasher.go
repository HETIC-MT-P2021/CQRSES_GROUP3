package services

import (
	"crypto/sha512"
	"fmt"
)

// HashPassword takes a string in parameter and returns the same string hashed with sha512
func HashPassword(password string) string {
	h := sha512.New()
	h.Write([]byte(password))
	bytesHash := h.Sum(nil)
	hexString := fmt.Sprintf("%x", bytesHash)
	// Returns the hexa string
	return hexString
}
