package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeToHash(presCode string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(presCode))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func hashCmp(presCode, hash string) bool {
	return hash == EncodeToHash(presCode)
}
