package stategen

import (
	"crypto/rand"
	"encoding/hex"
)

func Generate() (string, error) {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		// Should never be run, but check just in case
		return "", err
	}

	return hex.EncodeToString(b), nil
}
