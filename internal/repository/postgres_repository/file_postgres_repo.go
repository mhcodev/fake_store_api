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

func (p *PostgresFileRepository) UploadFile(ctx context.Context, file *models.File) error {
	query := `
		INSERT (
			filename,
			type,
			url
		) VALUES ($1, $2, $3)
		RETURNING id;
	`

	var fileID int

	err := p.conn.QueryRow(ctx, query,
		&file.FileName,
		&file.Type,
		&file.Url,
	).Scan(&fileID)

	if err != nil {
		return err
	}

	file.ID = fileID

	return nil
}
