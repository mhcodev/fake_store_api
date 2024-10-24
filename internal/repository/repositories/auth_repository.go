package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}
