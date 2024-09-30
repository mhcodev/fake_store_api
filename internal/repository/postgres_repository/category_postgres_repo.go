package postgresrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresCategoryRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresCategoryRepository(conn *pgxpool.Pool) *PostgresCategoryRepository {
	return &PostgresCategoryRepository{conn: conn}
}

// GetCategories returns a list of categories
func (p *PostgresCategoryRepository) GetCategories(ctx context.Context) ([]models.Category, error) {
	query := "SELECT id, name, image_url FROM tb_categories"

	rows, err := p.conn.Query(ctx, query)

	var categories []models.Category

	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var category models.Category
		rows.Scan(
			&category.ID,
			&category.Name,
			&category.ImageURL,
		)

		categories = append(categories, category)
	}

	return categories, nil
}
