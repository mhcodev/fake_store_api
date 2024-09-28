package postgresrepository

import (
	"context"
	"errors"
	"fmt"

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

	query := `SELECT
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
	from tb_users
	LIMIT $1
	OFFSET $2`

	rows, err := p.conn.Query(ctx, query, params.Limit, params.Offset)

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

func (p *PostgresUserRepository) GetUserByID(ctx context.Context, ID int) (models.User, error) {

	query := `SELECT
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
	FROM tb_users
	WHERE id = $1`

	row := p.conn.QueryRow(ctx, query, ID)
	var user models.User

	err := row.Scan(
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
		return user, err
	}

	return user, nil
}

func (p *PostgresUserRepository) UserEmailIsAvailable(ctx context.Context, email string) (bool, error) {

	query := `SELECT
		email
	FROM tb_users
	WHERE LOWER(email) = $1
	LIMIT 1`

	row := p.conn.QueryRow(ctx, query, email)
	var userEmail string

	err := row.Scan(&userEmail)

	fmt.Println("email: ", email)
	fmt.Println("userEmail: ", userEmail)
	fmt.Println("err: ", err)

	if err != nil {
		return true, nil
	}

	return false, nil
}

func (p *PostgresUserRepository) CreateUser(ctx context.Context, user *models.User) (bool, error) {
	query := `
	INSERT INTO tb_users (
			"user_type_id",
			"name",
			"email",
			"password",
			"avatar",
			"phone"
		)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id;`

	var userID int

	err := p.conn.QueryRow(ctx, query,
		&user.UserTypeID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Phone,
	).Scan(&userID)

	// Assign the new user id
	user.ID = userID

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *PostgresUserRepository) UpdateUser(ctx context.Context, user *models.User) (bool, error) {
	query := `UPDATE tb_users
	SET user_type_id = $1,
		name = $2,
		email = $3,
		password = $4,
		avatar = $5,
		phone = $6,
		updated_at = now()
	WHERE id = $7`

	commandTag, err := p.conn.Exec(ctx, query,
		&user.UserTypeID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Phone,
		&user.ID,
	)

	rowsAffected := commandTag.RowsAffected()

	if err != nil {
		return false, err
	}

	if rowsAffected <= 0 {
		return false, errors.New("no users were updated")
	}

	return true, nil
}
