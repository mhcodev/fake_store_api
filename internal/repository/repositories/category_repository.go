package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context) ([]models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
}
