package validators

import (
	"fmt"
	"regexp"
	"strings"
)

type ValidationErrors map[string][]string

// AddError adds an error message for a specific field.
func (ve ValidationErrors) AddError(field string, message string) {
	ve[field] = append(ve[field], message)
}

// HasErrors checks if any errors exist.
func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}

func (ve ValidationErrors) Error() string {
	return fmt.Sprintf("validation failed")
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

// IsValidEmail validates email format using a regex
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(strings.ToLower(email))
}

// IsNotEmpty checks if a string is not empty after trimming
func IsNotEmpty(str *string) bool {
	return str != nil && len(strings.TrimSpace(*str)) > 0
}

// IsStringLength checks if a string length is within the given range
func IsStringLength(str *string, min int, max int) bool {
	if str == nil {
		return false
	}
	length := len(strings.TrimSpace(*str))
	return length >= min && length <= max
}

// IsInRange checks if a number is within a specific range
func IsInRange(value *int, min int, max int) bool {
	if value == nil {
		return false
	}
	return *value >= min && *value <= max
}
