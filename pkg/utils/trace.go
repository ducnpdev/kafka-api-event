package utils

import (
	"encoding/hex"

	cryRand "crypto/rand"
)

func GenLogID() string {
	bytes := make([]byte, 8)
	_, err := cryRand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
