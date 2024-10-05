package util

import (
	"math/rand"
)

func GenerateRandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	sku := make([]byte, length)

	for i := range sku {
		sku[i] = charset[rand.Intn(len(charset))]
	}

	return string(sku)
}
