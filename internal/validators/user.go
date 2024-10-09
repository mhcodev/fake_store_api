package validators

import (
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/pkg"
)

// Validates the user input for create operation
func ValidateUserCreateInput(input services.UserCreateInput) ValidationErrors {
	var validationErrors = ValidationErrors{}

	if input.UserTypeID == nil &&
		input.Name == nil &&
		input.Email == nil &&
		input.Password == nil &&
		input.Avatar == nil &&
		input.Phone == nil &&
		input.Status == nil {
		validationErrors.AddError("error", "body request is not valid or empty")
		return validationErrors
	}

	if input.UserTypeID == nil {
		validationErrors.AddError("userTypeID", "userTypeID is required")
	} else if *input.UserTypeID <= 0 {
		validationErrors.AddError("userTypeID", "userTypeID must be a number grater than 0")
	}

	if input.Name == nil {
		validationErrors.AddError("name", "name is required")
	} else {
		nameSize := len(*input.Name)
		if IsNotEmpty(input.Name) && !IsInRange(&nameSize, 1, 36) {
			validationErrors.AddError("name", "name must be at least 1 letter and less than 36")
		}
	}

	if input.Email == nil {
		validationErrors.AddError("email", "email is required")
	} else if IsNotEmpty(input.Email) && !IsValidEmail(*input.Email) {
		validationErrors.AddError("email", "email is not valid")
	}

	if input.Password == nil || !IsNotEmpty(input.Password) {
		validationErrors.AddError("password", "password is required")
	} else {
		passSize := len(*input.Password)
		if !IsInRange(&passSize, 6, 21) {
			validationErrors.AddError("password", "password must be at least 6 characters up to 21 characters")
		}
	}

	if IsNotEmpty(input.Avatar) {
		if isImage, _ := pkg.IsImageURL(*input.Avatar); !isImage {
			validationErrors.AddError("avatar", "avatar is not a valid image")
		}
	}

	// if !IsNotEmpty(input.Phone) {
	// 	validationErrors.AddError("phone", "phone is not valid")
	// }

	return validationErrors
}

// Validates the user input for update operation
func ValidateUserUpdateInput(input services.UserUpdateInput) ValidationErrors {
	var validationErrors = ValidationErrors{}

	if input.UserTypeID == nil &&
		input.Name == nil &&
		input.Email == nil &&
		input.Password == nil &&
		input.Avatar == nil &&
		input.Phone == nil &&
		input.Status == nil {
		validationErrors.AddError("error", "body request is not valid or empty")
		return validationErrors
	}

	if input.UserTypeID != nil && (*input.UserTypeID <= 0) {
		validationErrors.AddError("userTypeID", "userTypeID must be a number grater than 0")
	}

	if !IsNotEmpty(input.Name) {
		validationErrors.AddError("name", "name must be at least 1 letter")
	}

	if !IsNotEmpty(input.Email) {
		validationErrors.AddError("email", "email is not valid")
	}

	if !IsNotEmpty(input.Password) {
		passSize := len(*input.Password)
		if !IsInRange(&passSize, 6, 21) {
			validationErrors.AddError("password", "password must be at least 6 characters up to 21 characters")
		}
	}

	if !IsNotEmpty(input.Avatar) {
		if isImage, _ := pkg.IsImageURL(*input.Avatar); !isImage {
			validationErrors.AddError("avatar", "avatar is not a valid image")
		}
	}

	// if !IsNotEmpty(input.Phone) {
	// 	validationErrors.AddError("phone", "phone is not valid")
	// }

	return validationErrors
}
