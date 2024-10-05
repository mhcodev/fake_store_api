package postgresrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresProductRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresProductRepository(conn *pgxpool.Pool) *PostgresProductRepository {
	return &PostgresProductRepository{conn: conn}
}

func (p *PostgresProductRepository) GetProductsByParams(ctx context.Context, params models.QueryParams) ([]models.Product, error) {
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
	`

	rows, err := p.conn.Query(ctx, query)

	var products []models.Product

	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
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
			return []models.Product{}, err
		}
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
		return product, err
	}

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
