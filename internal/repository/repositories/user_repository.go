package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type UserRepository interface {
	GetUsersByParams(ctx context.Context, params models.QueryParams) ([]models.User, error)
	GetUserByID(ctx context.Context, ID int) (models.User, error)
	UserEmailIsAvailable(ctx context.Context, email string) (bool, error)
}
