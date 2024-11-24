package services

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
	"github.com/mhcodev/fake_store_api/pkg"
	"golang.org/x/crypto/bcrypt"
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

type UserCreateInput struct {
	UserTypeID *int    `json:"userTypeID"`
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
	Avatar     *string `json:"avatar"`
	Phone      *string `json:"phone"`
	Status     *int8   `json:"status"`
}

func (s *UserService) CreateUser(ctx context.Context, input UserCreateInput) (*models.User, error) {
	isAvailable, _ := s.userRepository.UserEmailIsAvailable(ctx, *input.Email)

	if !isAvailable {
		return nil, errors.New("email is already used, try other email")
	}

	userTypes, err := s.GetUserTypes(ctx)

	if err != nil {
		return nil, errors.New("user types no available")
	}

	var typesAvailable []int

	for _, userType := range userTypes {
		typesAvailable = append(typesAvailable, userType.ID)
	}

	if len(typesAvailable) > 0 && !pkg.Includes(typesAvailable, *input.UserTypeID) {
		return nil, errors.New("user type id is not valid")
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(*input.Password), 10)

	if err != nil {
		return nil, errors.New("error generating password")
	}

	// Map input to user model
	newUser := &models.User{
		UserTypeID: *input.UserTypeID,
		Name:       *input.Name,
		Email:      *input.Email,
		Password:   string(passwordHashed),
		Avatar:     *input.Avatar,
	}

	if input.Avatar != nil {
		newUser.Avatar = *input.Avatar
	}

	if input.Phone != nil {
		newUser.Phone = *input.Phone
	}

	if input.Status != nil {
		newUser.Status = 1
	}

	ok, err := s.userRepository.CreateUser(ctx, newUser)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("user was not created, check your user data")
	}

	return newUser, nil
}

type UserUpdateInput struct {
	UserTypeID *int    `json:"userTypeID"`
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
	Avatar     *string `json:"avatar"`
	Phone      *string `json:"phone"`
	Status     *int8   `json:"status"`
}

func (s *UserService) UpdateUser(ctx context.Context, ID int, input UserUpdateInput) (*models.User, error) {
	user, err := s.GetUserByID(ctx, ID)

	if err != nil {
		return nil, errors.New("user not found")
	}

	userTypes, err := s.GetUserTypes(ctx)

	if err != nil {
		return nil, errors.New("user types no available")
	}

	var typesAvailable []int

	for _, userType := range userTypes {
		typesAvailable = append(typesAvailable, userType.ID)
	}

	if input.UserTypeID != nil {
		user.UserTypeID = *input.UserTypeID
	}

	if len(typesAvailable) > 0 && !pkg.Includes(typesAvailable, user.UserTypeID) {
		return nil, errors.New("user type id is not valid")
	}

	if input.Password != nil {
		passwordHashed, err := bcrypt.GenerateFromPassword([]byte(*input.Password), 10)

		if err != nil {
			return nil, errors.New("error generating password")
		}

		user.Password = string(passwordHashed)
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Email != nil {
		var emailIsAVailable bool

		if !strings.EqualFold(user.Email, *input.Email) {
			emailIsAVailable, _ = s.userRepository.UserEmailIsAvailable(ctx, *input.Email)
		} else {
			emailIsAVailable = true
		}

		if emailIsAVailable {
			user.Email = *input.Email
		} else {
			return nil, errors.New("email is already used by other user")
		}
	}

	if input.Avatar != nil {
		user.Avatar = *input.Avatar
	}

	if input.Phone != nil {
		user.Phone = *input.Phone
	}

	if input.Status != nil {
		user.Status = 1
	}

	ok, err := s.userRepository.UpdateUser(ctx, &user)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("user was not updated, check your user data")
	}

	return &user, nil
}

func (s *UserService) DeletedUser(ctx context.Context, userID int) error {
	ok, err := s.userRepository.DeleteUser(ctx, userID)

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("user was not deleted, check your user data")
	}

	return nil
}

func (s *UserService) GetUserTypes(ctx context.Context) ([]models.UserType, error) {
	userTypes, err := s.userRepository.GetUserTypes(ctx)
	return userTypes, err
}
