package validators

import "github.com/mhcodev/fake_store_api/internal/services"

// Validates LoginInput
func ValidateLoginInput(input services.LoginInput) ValidationErrors {
	var validationErrors = ValidationErrors{}

	if input.Email == nil &&
		input.Password == nil {
		validationErrors.AddError("error", "body request is not valid or empty")
		return validationErrors
	}

	if input.Email == nil {
		validationErrors.AddError("email", "email is required")
	} else if !IsNotEmpty(input.Email) {
		validationErrors.AddError("email", "email is empty")
	} else if !IsValidEmail(*input.Email) {
		validationErrors.AddError("email", "email is not valid")
	}

	if input.Password == nil {
		validationErrors.AddError("password", "password is required")
	} else if !IsNotEmpty(input.Password) {
		validationErrors.AddError("password", "password is empty")
	}

	return validationErrors
}

func ValidateNewTokenInput(input services.NewTokenInput) ValidationErrors {
	var validationErrors = ValidationErrors{}

	if input.RefreshToken == nil {
		validationErrors.AddError("error", "refreshToken is required")
		return validationErrors
	}

	if !IsNotEmpty(input.RefreshToken) {
		validationErrors.AddError("error", "refreshToken cannot be empty")
	}

	return validationErrors
}
