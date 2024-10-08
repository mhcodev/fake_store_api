package services

import (
	"context"
	"errors"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
	"github.com/mhcodev/fake_store_api/internal/validators"

	"github.com/mhcodev/fake_store_api/pkg"
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

func (ps *ProductService) CreateProduct(ctx context.Context, input models.ProductCreateInput) (*models.Product, error) {

	// Map input to product model
	newProduct := &models.Product{
		CategoryID:  *input.CategoryID,
		Name:        *input.Name,
		Slug:        validators.GenerateSlug(*input.Name),
		Description: *input.Description,
		Price:       *input.Price,
		Stock:       *input.Stock,
		Discount:    *input.Discount,
	}

	if input.Sku == nil {
		newProduct.Sku = pkg.GenerateRandomString(8)
	} else {
		newProduct.Sku = *input.Sku
	}

	err := ps.productRepository.CreateProduct(ctx, newProduct)

	if err != nil {
		return &models.Product{}, err
	}

	productUpdated, err := ps.productRepository.GetProductByID(ctx, newProduct.ID)

	if err != nil {
		return &models.Product{}, err
	}

	return &productUpdated, nil
}

func (ps *ProductService) UpdateProduct(ctx context.Context, ID int, input models.ProductUpdateInput) (*models.Product, error) {

	product, err := ps.GetProductByID(ctx, ID)

	if err != nil {
		return &models.Product{}, errors.New("product not found")
	}

	if input.CategoryID != nil {
		product.CategoryID = *input.CategoryID
	}

	if input.Name != nil {
		product.Name = *input.Name
		product.Slug = validators.GenerateSlug(*input.Name)
	}

	if input.Sku != nil {
		product.Sku = *input.Sku
	}

	if input.Description != nil {
		product.Description = *input.Description
	}

	if input.Price != nil {
		product.Price = *input.Price
	}

	if input.Stock != nil {
		product.Stock = *input.Stock
	}

	if input.Discount != nil {
		product.Discount = *input.Discount
	}

	err = ps.productRepository.UpdateProduct(ctx, &product)

	if err != nil {
		return &models.Product{}, errors.New("product no updated")
	}

	return &product, nil
}
