package hasher

import (
	"crypto/sha1"
	"encoding/hex"
)

func MakeSha(name string) string {
	// SHA1 hash
	h := sha1.New()
	h.Write([]byte(name))
	sha := h.Sum(nil)

	// Hexadecimal conversion
	shaStr := hex.EncodeToString(sha)
	return shaStr
}
