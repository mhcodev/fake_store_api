package validators

import (
	"fmt"
	"strings"

	"github.com/mhcodev/fake_store_api/internal/services"
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

// Validates the user input for update operation
func ValidateProductUpdateInput(input services.ProductUpdateInput) ValidationErrors {
	var validationErrors = ValidationErrors{}

	if input.CategoryID == nil &&
		input.Sku == nil &&
		input.Name == nil &&
		input.Stock == nil &&
		input.Description == nil &&
		input.Price == nil &&
		input.Discount == nil &&
		input.Status == nil {
		validationErrors.AddError("error", "body request is not valid or empty")
	}

	if input.CategoryID != nil && (*input.CategoryID <= 0) {
		validationErrors.AddError("categoryID", "categoryID must be a number grater than 0")
	}

	if input.Sku != nil && len(strings.TrimSpace(*input.Sku)) < 1 {
		validationErrors.AddError("sku", "sku must be at least 1 letter or number")
	}

	if input.Name != nil && strings.TrimSpace(*input.Name) == "" {
		validationErrors.AddError("name", "name is not valid")
	}

	if input.Stock != nil && (*input.Stock > 10000 || *input.Stock < 0) {
		validationErrors.AddError("stock", "Stock must be between 0 and 10000")
	}

	if input.Description != nil && strings.TrimSpace(*input.Description) == "" {
		validationErrors.AddError("description", "description is not valid")
	}

	if input.Price != nil && (*input.Price > 1000000 || *input.Price < 0) {
		validationErrors.AddError("price", "price must be grater or equal than 0")
	}

	if input.Discount != nil && (*input.Discount < 0 || *input.Discount > 1) {
		validationErrors.AddError("discount", "name is not valid")
	}

	if input.Status != nil && (*input.Status < 0 || *input.Status > 100) {
		validationErrors.AddError("status", "status must be between 0 and 100")
	}

	return validationErrors
}

// Simple email validation function (you can improve the regex as needed)
func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
