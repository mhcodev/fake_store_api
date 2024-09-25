package services

import (
	"context"

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
