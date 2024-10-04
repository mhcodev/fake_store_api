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
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.CategoryID,
			&p.Sku,
			&p.Name,
			&p.Slug,
			&p.Stock,
			&p.Description,
			&p.Price,
			&p.Discount,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return []models.Product{}, err
		}
		products = append(products, p)
	}

	return products, nil
}
