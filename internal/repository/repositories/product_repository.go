package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type ProductRepository interface {
	GetProductsByParams(ctx context.Context, params models.QueryParams) ([]models.Product, error)
	GetProductByID(ctx context.Context, ID int) (models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) error
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, ID int) error
}
