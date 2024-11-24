package pkg

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
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

// Includes checks if a number is in the slice
func Includes(arr []int, num int) bool {
	for _, value := range arr {
		if value == num {
			return true
		}
	}
	return false
}

// IsImageURL checks if the Content-Type starts with "image/"
func IsImageURL(url string) (bool, error) {
	resp, err := http.Head(url)
	if err != nil {
		return false, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return false, fmt.Errorf("could not find Content-Type header")
	}

	// Check if the Content-Type starts with "image/"
	return strings.HasPrefix(contentType, "image/"), nil
}

// GetBaseURL return the Base URL from a link as a parameter
func GetBaseURL(link string) (string, error) {
	// Parse the input URL
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	// Extract the base URL (scheme + host + optional port)
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	return baseURL, nil
}
