package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type UserRepository interface {
	GetTotalUsers(ctx context.Context) (int, error)
	GetUsersByParams(ctx context.Context, params models.QueryParams) ([]models.User, error)
	GetUserByID(ctx context.Context, ID int) (models.User, error)
	UserEmailIsAvailable(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *models.User) (bool, error)
	UpdateUser(ctx context.Context, user *models.User) (bool, error)
	DeleteUser(ctx context.Context, userID int) (bool, error)
	GetUserTypes(ctx context.Context) ([]models.UserType, error)
}
