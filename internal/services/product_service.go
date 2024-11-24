package services

import (
	"context"
	"errors"
	"sync"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"

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

func (ps *ProductService) GetTotalOfProducts(ctx context.Context) (int, error) {
	count, err := ps.productRepository.GetTotalOfProducts(ctx)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ps *ProductService) GetProductsByParams(ctx context.Context, params models.QueryParams) ([]models.Product, error) {
	// Checks limit & offset default value
	if params.Limit < 1 {
		params.Limit = 15
	}

	if params.Offset < 0 {
		params.Offset = 0
	}

	products, err := ps.productRepository.GetProductsByParams(ctx, params)

	if err != nil {
		return []models.Product{}, err
	}

	var wg sync.WaitGroup

	for idx, product := range products {
		wg.Add(1)
		go func() {
			defer wg.Done()

			images, _ := ps.productRepository.GetImagesByProduct(ctx, product.ID)
			urls := make([]string, 0)

			for _, image := range images {
				urls = append(urls, image.ImageURL)
			}

			product.Images = urls
			products[idx] = product
		}()
	}

	wg.Wait()

	return products, nil
}

func (ps *ProductService) GetProductByID(ctx context.Context, ID int) (models.Product, error) {

	product, err := ps.productRepository.GetProductByID(ctx, ID)

	if err != nil {
		return models.Product{}, err
	}

	images, _ := ps.productRepository.GetImagesByProduct(ctx, ID)

	urls := make([]string, 0)

	for _, image := range images {
		urls = append(urls, image.ImageURL)
	}

	product.Images = urls

	return product, nil
}

type ProductCreateInput struct {
	CategoryID  *int      `json:"categoryID"`
	Sku         *string   `json:"sku"`
	Name        *string   `json:"name"`
	Stock       *int      `json:"stock"`
	Description *string   `json:"description"`
	Price       *float32  `json:"price"`
	Images      *[]string `json:"images"`
	Discount    *float32  `json:"discount"`
	Status      *int8     `json:"status"`
}

func (ps *ProductService) CreateProduct(ctx context.Context, input ProductCreateInput) (*models.Product, error) {
	// Map input to product model
	newProduct := &models.Product{
		CategoryID:  *input.CategoryID,
		Name:        *input.Name,
		Slug:        pkg.GenerateSlug(*input.Name),
		Description: *input.Description,
		Price:       *input.Price,
		Stock:       0,
		Discount:    0,
	}

	if input.Sku == nil {
		newProduct.Sku = pkg.GenerateRandomString(8)
	} else {
		skuIsAvailable, _ := ps.productRepository.SkuIsAvailable(ctx, *input.Sku)

		if skuIsAvailable {
			newProduct.Sku = *input.Sku
		} else {
			return &models.Product{}, errors.New("sku is not available")
		}
	}

	if input.Stock != nil {
		newProduct.Stock = *input.Stock
	}

	if input.Discount != nil {
		newProduct.Discount = *input.Discount
	}

	err := ps.productRepository.CreateProduct(ctx, newProduct)

	if err != nil {
		return &models.Product{}, err
	}

	productUpdated, err := ps.productRepository.GetProductByID(ctx, newProduct.ID)

	if err != nil {
		return &models.Product{}, err
	}

	if input.Images != nil {
		validImages, _ := ps.productRepository.AssiociateImagesToProduct(ctx, productUpdated.ID, *input.Images)
		productUpdated.Images = validImages
	}

	return &productUpdated, nil
}

type ProductUpdateInput struct {
	CategoryID  *int      `json:"categoryID"`
	Sku         *string   `json:"sku"`
	Name        *string   `json:"name"`
	Stock       *int      `json:"stock"`
	Description *string   `json:"description"`
	Price       *float32  `json:"price"`
	Discount    *float32  `json:"discount"`
	Status      *int8     `json:"status"`
	Images      *[]string `json:"images"`
}

func (ps *ProductService) UpdateProduct(ctx context.Context, ID int, input ProductUpdateInput) (*models.Product, error) {

	product, err := ps.GetProductByID(ctx, ID)

	if err != nil {
		return &models.Product{}, errors.New("product not found")
	}

	if input.CategoryID != nil {
		product.CategoryID = *input.CategoryID
	}

	if input.Name != nil {
		product.Name = *input.Name
		product.Slug = pkg.GenerateSlug(*input.Name)
	}

	if input.Sku != nil {
		var skuIsAvailable bool

		if product.Sku != *input.Sku {
			skuIsAvailable, _ = ps.productRepository.SkuIsAvailable(ctx, *input.Sku)
		} else {
			skuIsAvailable = true
		}

		if skuIsAvailable {
			product.Sku = *input.Sku
		} else {
			return &models.Product{}, errors.New("sku is not available")
		}
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

	if input.Images != nil {
		ps.productRepository.DeleteImagesByProduct(ctx, product.ID)
		validImages, _ := ps.productRepository.AssiociateImagesToProduct(ctx, product.ID, *input.Images)
		product.Images = validImages
	}

	err = ps.productRepository.UpdateProduct(ctx, &product)

	if err != nil {
		return &models.Product{}, errors.New("product no updated")
	}

	return &product, nil
}

func (ps *ProductService) DeleteProduct(ctx context.Context, ID int) error {
	_, err := ps.GetProductByID(ctx, ID)

	if err != nil {
		return errors.New("product not found")
	}

	err = ps.productRepository.DeleteProduct(ctx, ID)

	if err != nil {
		return errors.New(err.Error())
	}

	err = ps.productRepository.DeleteImagesByProduct(ctx, ID)

	if err != nil {
		return errors.New("images were no deleted")
	}

	return nil
}
