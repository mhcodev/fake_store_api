package postgresrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresUserRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresUserRepository(conn *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{conn: conn}
}

func (p *PostgresUserRepository) GetUsersByParams(ctx context.Context, params models.QueryParams) ([]models.User, error) {

	query := `select
		id,
		user_type_id,
		name,
		email,
		password,
		avatar,
		phone,
		status,
		created_at,
		updated_at
	from tb_users`

	rows, err := p.conn.Query(context.Background(), query)

	if err != nil {
		return []models.User{}, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User

		err := rows.Scan(
			&u.ID,
			&u.UserTypeID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.Avatar,
			&u.Phone,
			&u.Status,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, u)
	}

	return users, nil
}
