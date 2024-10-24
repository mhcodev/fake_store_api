package postgresrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresAuthRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresAuthRepository(conn *pgxpool.Pool) *PostgresAuthRepository {
	return &PostgresAuthRepository{conn: conn}
}

// GetUserByEmail returns an user by email
func (p *PostgresAuthRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := `
		SELECT  id,
				user_type_id,
			    name,
				email,
				password,
				avatar,
				phone,
				status,
				created_at,
				updated_at
		FROM tb_users
		WHERE email = $1 and status = 1
	`

	var user models.User

	err := p.conn.QueryRow(ctx, query,
		&email,
	).Scan(
		&user.ID,
		&user.UserTypeID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Phone,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
