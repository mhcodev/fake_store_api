package services

import (
	"context"
	"errors"
	"regexp"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: *userRepository,
	}
}

func (s *UserService) GetUsersByParams(ctx context.Context, params models.QueryParams) ([]models.User, error) {
	if params.Limit < 1 {
		params.Limit = 15
	}

	if params.Offset < 0 {
		params.Offset = 0
	}

	return s.userRepository.GetUsersByParams(ctx, params)
}

func (s *UserService) GetUserByID(ctx context.Context, ID int) (models.User, error) {
	if ID < 1 {
		return models.User{}, errors.New("user id is not valid")
	}

	return s.userRepository.GetUserByID(ctx, ID)
}

func (s *UserService) UserEmailIsAvailable(ctx context.Context, email string) (map[string]interface{}, error) {
	// Define a regular expression for validating email addresses
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	rgx := regexp.MustCompile(emailRegex)

	// Check if the email matches the pattern
	emailIsValid := rgx.MatchString(email)

	response := make(map[string]interface{})

	if !emailIsValid {
		response["isAvailable"] = false
		response["formatIsOK"] = false
		return response, errors.New("email provided is not valid")
	}

	IsVailable, err := s.userRepository.UserEmailIsAvailable(ctx, email)

	if err != nil {
		response["isAvailable"] = false
		response["formatIsOK"] = false
		response["error"] = err.Error()
		return response, err
	}

	response["isAvailable"] = IsVailable
	response["formatIsOK"] = true

	return response, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	ok, err := s.userRepository.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("user was not created, check your user data")
	}

	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	ok, err := s.userRepository.UpdateUser(ctx, user)

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("user was not updated, check your user data")
	}

	return nil
}
