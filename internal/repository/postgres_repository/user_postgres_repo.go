package postgresrepository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresUserRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresUserRepository(conn *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{conn: conn}
}

// GetUsersByParams search a list of users by params
// @params - models.QueryParams
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
	from tb_users`

	mapParams := params.MapParams

	var queryParams string

	if name := mapParams["name"]; name != "" {
		queryParams = fmt.Sprintf("LOWER(name)='%s'", strings.ToLower(name.(string)))
	} else if userTypeID := mapParams["type"]; userTypeID != "" && userTypeID != -1 {
		queryParams = fmt.Sprintf("user_type_id=%d", userTypeID)
	} else if email := mapParams["email"]; email != "" {
		queryParams = fmt.Sprintf("email='%s'", email)
	} else if status := mapParams["status"]; status != "" && status != -1 {
		queryParams = fmt.Sprintf("status=%d", status)
	}

	if queryParams != "" {
		query += " WHERE "
		query += fmt.Sprintf("%s ", queryParams)
		query += " LIMIT $1 OFFSET $2"
	} else {
		query += " LIMIT $1 OFFSET $2"
	}

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

// GetUserByID returns a user by id
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
	WHERE id = $1 and status = 1`

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

// UserEmailIsAvailable checks an user email is available
func (p *PostgresUserRepository) UserEmailIsAvailable(ctx context.Context, email string) (bool, error) {

	query := `SELECT
		email
	FROM tb_users
	WHERE LOWER(email) = $1
	LIMIT 1`

	row := p.conn.QueryRow(ctx, query, strings.ToLower(email))
	var userEmail string

	err := row.Scan(&userEmail)

	if err != nil {
		return true, nil
	}

	return false, nil
}

// CreateUser creates a user to db
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

// UpdateUser updates a user by id
func (p *PostgresUserRepository) UpdateUser(ctx context.Context, user *models.User) (bool, error) {
	query := `UPDATE tb_users
	SET user_type_id = $1,
		name = $2,
		email = $3,
		password = $4,
		avatar = $5,
		phone = $6,
		status = $7,
		updated_at = now()
	WHERE id = $8`

	commandTag, err := p.conn.Exec(ctx, query,
		&user.UserTypeID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Phone,
		&user.Status,
		&user.ID,
	)

	rowsAffected := commandTag.RowsAffected()

	if err != nil {
		fmt.Println("error:", err.Error())
		return false, errors.New("no users were updated")
	}

	if rowsAffected <= 0 {
		fmt.Println("error:", "No rows affected")
		return false, errors.New("no users were updated")
	}

	return true, nil
}

// DeleteUser deletes a user by id (status = 0)
func (p *PostgresUserRepository) DeleteUser(ctx context.Context, userID int) (bool, error) {
	query := `UPDATE tb_users
	SET status = $1,
		updated_at = now()
	WHERE id = $2`

	statusDeletedCode := 0

	commandTag, err := p.conn.Exec(ctx, query,
		statusDeletedCode,
		userID,
	)

	rowsAffected := commandTag.RowsAffected()

	if err != nil {
		return false, err
	}

	if rowsAffected <= 0 {
		return false, errors.New("no users were deleted")
	}

	return true, nil
}

// GetUserTypes returns a list of user types
func (p *PostgresUserRepository) GetUserTypes(ctx context.Context) ([]models.UserType, error) {
	query := "SELECT id, name, description FROM tb_user_type"

	rows, err := p.conn.Query(ctx, query)

	var userTypes []models.UserType

	if err != nil {
		return userTypes, err
	}

	for rows.Next() {
		var userType models.UserType
		rows.Scan(
			&userType.ID,
			&userType.Name,
			&userType.Description,
		)

		userTypes = append(userTypes, userType)
	}

	return userTypes, nil
}
