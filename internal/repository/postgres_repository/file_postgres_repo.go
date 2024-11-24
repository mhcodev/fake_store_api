package postgresrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresFileRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresFileRepository(conn *pgxpool.Pool) *PostgresFileRepository {
	return &PostgresFileRepository{conn: conn}
}

func (p *PostgresFileRepository) SaveFileToDB(ctx context.Context, file *models.File) error {
	query := `
		INSERT INTO tb_files (
			original_name,
			filename,
			type,
			url,
			base_url
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	var fileID int

	err := p.conn.QueryRow(ctx, query,
		&file.OriginalName,
		&file.FileName,
		&file.Type,
		&file.Url,
		&file.BaseURL,
	).Scan(&fileID)

	if err != nil {
		return err
	}

	file.ID = fileID

	return nil
}
