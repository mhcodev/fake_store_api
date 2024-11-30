package postgresrepository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/pkg"
)

type PostgresProductRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresProductRepository(conn *pgxpool.Pool) *PostgresProductRepository {
	return &PostgresProductRepository{conn: conn}
}

// GetTotalOfProducts returns the total quantity of active products
func (p *PostgresProductRepository) GetTotalOfProducts(ctx context.Context) (int, error) {
	query := `
		SELECT COUNT(*) as total
		FROM tb_products
		WHERE status = 1;
	`

	var count int

	err := p.conn.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		return count, err
	}

	return count, nil
}

// SkuIsAvailable checks if product sku is available (duplication is not allowed)
func (p *PostgresProductRepository) SkuIsAvailable(ctx context.Context, sku string) (bool, error) {
	query := `
		SELECT COUNT(*) as total
		FROM tb_products
		WHERE UPPER(sku) = $1 AND status = 1;
	`

	var count int

	err := p.conn.QueryRow(ctx, query,
		strings.ToUpper(sku),
	).Scan(&count)

	if err != nil {
		return false, err
	}

	isAvailable := count == 0

	return isAvailable, nil
}

func (p *PostgresProductRepository) GetProductsByParams(ctx context.Context, params models.QueryParams) ([]models.Product, error) {
	query := `
		SELECT
			p.id,
			p.category_id,
			p.sku,
			p.name,
			p.slug,
			p.stock,
			p.description,
			p.price,
			p.discount,
			p.status,
			p.created_at,
			p.updated_at,
			c.id,
			c.name,
			c.image_url,
			c.status
		FROM tb_products as p
		INNER JOIN tb_categories c
		ON c.id = p.category_id
		LIMIT $1
		OFFSET $2;
	`

	rows, err := p.conn.Query(ctx, query,
		&params.Limit,
		&params.Offset,
	)

	products := make([]models.Product, 0)

	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		var category models.Category

		err = rows.Scan(
			&product.ID,
			&product.CategoryID,
			&product.Sku,
			&product.Name,
			&product.Slug,
			&product.Stock,
			&product.Description,
			&product.Price,
			&product.Discount,
			&product.Status,
			&product.CreatedAt,
			&product.UpdatedAt,
			&category.ID,
			&category.Name,
			&category.ImageURL,
			&category.Status,
		)

		if err != nil {
			return []models.Product{}, err
		}

		product.Category = category

		products = append(products, product)
	}

	return products, nil
}

// GetProductByID returns a product by id
func (p *PostgresProductRepository) GetProductByID(ctx context.Context, ID int) (models.Product, error) {
	query := `
		SELECT
			id,
			category_id,
			sku,
			name,
			slug,
			stock,
			description,
			price,
			discount,
			status,
			created_at,
			updated_at
		FROM tb_products
		WHERE id = $1
	`

	var product models.Product

	err := p.conn.QueryRow(ctx, query, ID).Scan(
		&product.ID,
		&product.CategoryID,
		&product.Sku,
		&product.Name,
		&product.Slug,
		&product.Stock,
		&product.Description,
		&product.Price,
		&product.Discount,
		&product.Status,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p *PostgresProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	query := `
		INSERT INTO
			tb_products (
				"category_id",
				"sku",
				"name",
				"slug",
				"stock",
				"description",
				"price",
				"discount"
			)
		VALUES
			(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8
			)
		RETURNING id;
	`

	var productID int

	err := p.conn.QueryRow(ctx, query,
		&product.CategoryID,
		&product.Sku,
		&product.Name,
		&product.Slug,
		&product.Stock,
		&product.Description,
		&product.Price,
		&product.Discount,
	).Scan(&productID)

	if err != nil {
		return err
	}

	product.ID = productID

	return nil
}

func (p *PostgresProductRepository) UpdateProduct(ctx context.Context, product *models.Product) error {
	query := `
		UPDATE tb_products
		SET category_id = $1,
			sku = $2,
			name = $3,
			slug = $4,
			stock = $5,
			description = $6,
			price = $7,
			discount = $8,
			updated_at = now()
		WHERE id = $9;
	`

	commandTag, err := p.conn.Exec(ctx, query,
		&product.CategoryID,
		&product.Sku,
		&product.Name,
		&product.Slug,
		&product.Stock,
		&product.Description,
		&product.Price,
		&product.Discount,
		&product.ID,
	)

	rowsAffected := commandTag.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("products were not updated")
	}

	return nil
}

func (p *PostgresProductRepository) DeleteProduct(ctx context.Context, ID int) error {
	query := `
		UPDATE tb_products
		SET status = $1,
			updated_at = now()
		WHERE id = $2;
	`

	statusDeletedCode := 0

	commandTag, err := p.conn.Exec(ctx, query,
		statusDeletedCode,
		ID,
	)

	rowsAffected := commandTag.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("products were not deleted")
	}

	return nil
}

func (p *PostgresProductRepository) AssiociateImagesToProduct(ctx context.Context, prodID int, urls []string) ([]string, []string) {
	query := `
		INSERT INTO tb_product_images (
			product_id,
			image_url,
			base_url
		) VALUES ($1, $2, $3);
	`

	validImages := make([]string, 0)
	errors := make([]string, 0)

	if len(urls) <= 0 {
		errors = append(errors, "image urls not found")
		return validImages, errors
	}

	for _, url := range urls {
		baseURL, err := pkg.GetBaseURL(url)

		if err != nil {
			continue
		}

		_, err = p.conn.Exec(ctx, query,
			&prodID,
			&url,
			&baseURL,
		)

		if err != nil {
			fmt.Println("error associating image to product")
			errors = append(errors, err.Error())
			continue
		}

		validImages = append(validImages, url)
	}

	return validImages, errors
}

func (p *PostgresProductRepository) GetImagesByProduct(ctx context.Context, prodID int) ([]models.ProductImage, error) {
	query := `
		SELECT id,
			   product_id,
			   image_url,
			   status,
			   created_at,
			   updated_at
		FROM tb_product_images
		WHERE product_id = $1 and status = 1;
	`

	rows, err := p.conn.Query(ctx, query, prodID)

	if err != nil {
		return []models.ProductImage{}, err
	}

	defer rows.Close()

	images := make([]models.ProductImage, 0)

	for rows.Next() {
		var prodImage models.ProductImage

		err = rows.Scan(
			&prodImage.ID,
			&prodImage.ProductID,
			&prodImage.ImageURL,
			&prodImage.Status,
			&prodImage.CreatedAt,
			&prodImage.UpdatedAt,
		)

		if err != nil {
			return []models.ProductImage{}, err
		}

		images = append(images, prodImage)
	}

	return images, nil
}

func (p *PostgresProductRepository) GetImagesByProducListID(ctx context.Context, IDs []int) ([]models.ProductImage, error) {
	query := `
		SELECT id,
			   product_id,
			   image_url,
			   status,
			   created_at,
			   updated_at
		FROM tb_product_images
		WHERE product_id = ANY ($1);
	`

	rows, err := p.conn.Query(ctx, query, IDs)

	if err != nil {
		return []models.ProductImage{}, err
	}

	defer rows.Close()

	images := make([]models.ProductImage, 0)

	for rows.Next() {
		var prodImage models.ProductImage

		err = rows.Scan(
			&prodImage.ID,
			&prodImage.ProductID,
			&prodImage.ImageURL,
			&prodImage.Status,
			&prodImage.CreatedAt,
			&prodImage.UpdatedAt,
		)

		if err != nil {
			return []models.ProductImage{}, err
		}

		images = append(images, prodImage)
	}

	return images, nil
}

func (p *PostgresProductRepository) DeleteImagesByProduct(ctx context.Context, prodID int) error {
	query := `
		UPDATE tb_product_images
		SET status = $1,
			updated_at = now()
		WHERE id = $2 and status = 1;
	`

	statusDeletedCode := 0

	commandTag, err := p.conn.Exec(ctx, query,
		statusDeletedCode,
		prodID,
	)

	rowsAffected := commandTag.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("images were not deleted")
	}

	return nil
}
