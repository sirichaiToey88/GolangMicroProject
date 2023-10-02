package other

import (
	"crypto/rand"
	"math/big"
)

func RandomStringWithNumbers(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	maxIdx := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		idx, _ := rand.Int(rand.Reader, maxIdx)
		result[i] = charset[idx.Int64()]
	}

	return string(result)
}
