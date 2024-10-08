package pkg

import (
	"math/rand"
	"regexp"
	"strings"
)

func GenerateRandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	sku := make([]byte, length)

	for i := range sku {
		sku[i] = charset[rand.Intn(len(charset))]
	}

	return string(sku)
}

// GenerateSlug converts to alphanumeric slug concated with dashes
func GenerateSlug(text string) string {
	slug := strings.ToLower(text)
	slug = strings.ReplaceAll(slug, " ", "-")

	re := regexp.MustCompile(`[^a-z0-9\-]+`)
	slug = re.ReplaceAllString(slug, "")

	slug = strings.Trim(slug, "-")
	return slug
}
