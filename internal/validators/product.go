package validators

import (
	"fmt"
	"strings"

	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/pkg"
)

// Validates the product input for create operation
func ValidateProductCreateInput(input services.ProductCreateInput) ValidationErrors {
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
		return validationErrors
	}

	if input.CategoryID == nil || (*input.CategoryID <= 0) {
		validationErrors.AddError("categoryID", "categoryID must be a number grater than 0")
	}

	if input.Sku != nil && len(strings.TrimSpace(*input.Sku)) < 1 {
		validationErrors.AddError("sku", "sku must be at least 1 letter or number")
	}

	if input.Name == nil || strings.TrimSpace(*input.Name) == "" {
		validationErrors.AddError("name", "name is required")
	}

	if input.Stock != nil && (*input.Stock > 10000 || *input.Stock < 0) {
		validationErrors.AddError("stock", "Stock must be between 0 and 10000")
	}

	if input.Description == nil || strings.TrimSpace(*input.Description) == "" {
		validationErrors.AddError("description", "description is not valid")
	}

	if input.Price == nil || (*input.Price > 1000000 || *input.Price < 0) {
		validationErrors.AddError("price", "price must be grater or equal than 0")
	}

	if input.Discount != nil && (*input.Discount < 0 || *input.Discount > 1) {
		validationErrors.AddError("discount", "name is not valid")
	}

	if input.Status != nil && (*input.Status < 0 || *input.Status > 100) {
		validationErrors.AddError("status", "status must be between 0 and 100")
	}

	if input.Images != nil {
		if len(*input.Images) > 6 {
			validationErrors.AddError("images", "max. 6 images per product")
		} else {
			for _, imageURL := range *input.Images {
				isValid, err := pkg.IsImageURL(imageURL)

				if err != nil || !isValid {
					msg := fmt.Sprintf("%s is not a valid image url", imageURL)
					validationErrors.AddError("images", msg)
				}
			}
		}
	}

	return validationErrors
}

// Validates the product input for update operation
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
		return validationErrors
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

	if input.Images != nil {
		if len(*input.Images) > 6 {
			validationErrors.AddError("images", "max. 6 images per product")
		} else {
			for _, imageURL := range *input.Images {
				isValid, err := pkg.IsImageURL(imageURL)

				if err != nil || !isValid {
					msg := fmt.Sprintf("%s is not a valid image url", imageURL)
					validationErrors.AddError("images", msg)
				}
			}
		}
	}

	return validationErrors
}
