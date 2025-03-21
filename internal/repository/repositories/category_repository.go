package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type CategoryRepository interface {
	GetTotalOfCategories(ctx context.Context) (int, error)
	GetCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, ID int) (models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
	UpdateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, ID int) error
}
