package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type ProductRepository interface {
	GetTotalOfProducts(ctx context.Context) (int, error)
	GetProductsByParams(ctx context.Context, params models.QueryParams) ([]models.Product, error)
	GetProductByID(ctx context.Context, ID int) (models.Product, error)
	SkuIsAvailable(ctx context.Context, sku string) (bool, error)
	CreateProduct(ctx context.Context, product *models.Product) error
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, ID int) error
	AssiociateImagesToProduct(ctx context.Context, prodID int, urls []string) ([]string, []string)
	GetImagesByProduct(ctx context.Context, prodID int) ([]models.ProductImage, error)
	DeleteImagesByProduct(ctx context.Context, prodID int) error
}
