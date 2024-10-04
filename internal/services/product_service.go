package services

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: *productRepository,
	}
}

func (ps *ProductService) GetProductsByParams(ctx context.Context, params models.QueryParams) ([]models.Product, error) {
	return ps.productRepository.GetProductsByParams(ctx, params)
}
