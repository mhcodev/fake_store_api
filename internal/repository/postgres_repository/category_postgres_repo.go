package postgresrepository

import (
	"context"
	"errors"

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

// GetCategoryByID returns a single category by ID
func (p *PostgresCategoryRepository) GetCategoryByID(ctx context.Context, ID int) (models.Category, error) {
	query := `
		SELECT id, name, image_url, status
		FROM tb_categories
		WHERE status = 1 AND id = $1
	`

	var category models.Category

	err := p.conn.QueryRow(ctx, query, ID).Scan(
		&category.ID,
		&category.Name,
		&category.ImageURL,
		&category.Status,
	)

	if err != nil {
		return category, err
	}

	return category, nil
}

// CreateCategory creates a category into db
func (p *PostgresCategoryRepository) CreateCategory(ctx context.Context, category *models.Category) error {
	query := `		
		INSERT INTO tb_categories ("name", "image_url")
		VALUES ($1, $2)
		RETURNING id;
	`

	var categoryID int

	err := p.conn.QueryRow(ctx, query,
		&category.Name,
		&category.ImageURL,
	).Scan(&categoryID)

	if err != nil {
		return err
	}

	category.ID = categoryID
	category.Status = 1

	return nil
}

// UpdateCategory updates a category into db
func (p *PostgresCategoryRepository) UpdateCategory(ctx context.Context, category *models.Category) error {
	query := `
		UPDATE tb_categories
		SET name = $1, image_url = $2, status = $3
		WHERE id = $4;
	`

	commandTag, err := p.conn.Exec(ctx, query,
		&category.Name,
		&category.ImageURL,
		&category.Status,
		&category.ID,
	)

	rowsAffected := commandTag.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("no categories were updated")
	}

	return nil
}

// DeleteCategory deletes a category by id (status = 0)
func (p *PostgresCategoryRepository) DeleteCategory(ctx context.Context, ID int) error {
	query := `
		UPDATE tb_categories
		SET status = $1
		WHERE id = $2;
	`

	statusDeletedCode := 0

	commandTag, err := p.conn.Exec(ctx, query,
		&statusDeletedCode,
		&ID,
	)

	rowsAffected := commandTag.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("no categories were updated")
	}

	return nil
}
