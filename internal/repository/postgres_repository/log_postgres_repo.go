package postgresrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresLogRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresLogRepository(conn *pgxpool.Pool) *PostgresLogRepository {
	return &PostgresLogRepository{conn: conn}
}

func (p *PostgresLogRepository) InsertApiLog(ctx context.Context, apiLog *models.ApiLog) error {
	query := `
		INSERT INTO tb_api_logs (method, version, path, status_code, response_time, user_id, ip_address, country)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	var logID int

	err := p.conn.QueryRow(ctx, query,
		&apiLog.Method,
		&apiLog.Version,
		&apiLog.Path,
		&apiLog.StatusCode,
		&apiLog.ResponseTime,
		&apiLog.UserID,
		&apiLog.IPAdress,
		&apiLog.Country,
	).Scan(&logID)

	if err != nil {
		return err
	}

	apiLog.ID = logID

	return nil
}
