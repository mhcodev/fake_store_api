package models

import (
	"regexp"
	"strings"
)

func (p *Product) Validate() []string {
	var messages []string

	if strings.TrimSpace(p.Name) == "" {
		messages = append(messages, "name is required")
	}

	if p.Discount < 0 || p.Discount > 1 {
		messages = append(messages, "discount must be grater than 0 and less that 1")
	}

	if p.Stock < 0 || p.Stock > 100000 {
		messages = append(messages, "stock must be equal or grater that 0 and less than 100000")
	}

	if strings.TrimSpace(p.Description) == "" {
		messages = append(messages, "description is required")
	}

	if !(p.Price >= 0 && p.Price <= 1000000.0) {
		messages = append(messages, "price must be equal or grater than 0 and less than 1000000.0")
	}

	if p.Status < 0 {
		messages = append(messages, "status must be equal or grater than 0")
	}

	return messages
}

func GenerateSlug(text string) string {
	slug := strings.ToLower(text)
	slug = strings.ReplaceAll(slug, " ", "-")

	re := regexp.MustCompile(`[^a-z0-9\-]+`)
	slug = re.ReplaceAllString(slug, "")

	slug = strings.Trim(slug, "-")
	return slug
}
