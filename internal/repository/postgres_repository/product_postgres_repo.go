package postgresrepository

import (
	"context"
	"errors"

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
		SET status = $1
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
