package services

import (
	"context"
	"errors"

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

func (ps *ProductService) GetProductByID(ctx context.Context, ID int) (models.Product, error) {
	return ps.productRepository.GetProductByID(ctx, ID)
}

func (ps *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	product.Slug = models.GenerateSlug(product.Name)
	err := ps.productRepository.CreateProduct(ctx, product)

	if err != nil {
		return err
	}

	productUpdated, err := ps.productRepository.GetProductByID(ctx, product.ID)

	if err != nil {
		return err
	}

	product.ID = productUpdated.ID
	product.CategoryID = productUpdated.CategoryID
	product.Sku = productUpdated.Sku
	product.Name = productUpdated.Name
	product.Slug = productUpdated.Slug
	product.Stock = productUpdated.Stock
	product.Description = productUpdated.Description
	product.Price = productUpdated.Price
	product.Discount = productUpdated.Discount
	product.Status = productUpdated.Status
	product.CreatedAt = productUpdated.CreatedAt
	product.UpdatedAt = productUpdated.UpdatedAt

	return nil
}

type ProductUpdateInput struct {
	CategoryID  *int     `json:"categoryID"`
	Sku         *string  `json:"sku"`
	Name        *string  `json:"name"`
	Stock       *int     `json:"stock"`
	Description *string  `json:"description"`
	Price       *float32 `json:"price"`
	Discount    *float32 `json:"discount"`
	Status      *int8    `json:"status"`
}

func (ps *ProductService) UpdateProduct(ctx context.Context, ID int, input ProductUpdateInput) (*models.Product, error) {

	product, err := ps.GetProductByID(ctx, ID)

	if err != nil {
		return &models.Product{}, errors.New("product not found")
	}

	err = ps.productRepository.UpdateProduct(ctx, &product)

	if err != nil {
		return &models.Product{}, errors.New("product no updated")
	}

	return &product, nil
}
