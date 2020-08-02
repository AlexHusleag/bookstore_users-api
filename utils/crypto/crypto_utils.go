package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSHA(input string) string {
	hash := sha256.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
