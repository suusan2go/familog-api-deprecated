package tokenGenerator

import (
	"crypto/rand"
	"encoding/base32"
)

// GenerateRandomToken generates token base32
func GenerateRandomToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
