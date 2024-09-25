package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/driver"
	postgresrepo "github.com/mhcodev/fake_store_api/internal/repository/postgres_repo"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type DBRepository struct {
	UserRepository repositories.UserRepository
}

func NewDBRepository(dbRepo *DBRepository) *DBRepository {
	return &DBRepository{
		UserRepository: dbRepo.UserRepository,
	}
}

func InitPosgresRepositories() (*DBRepository, *pgxpool.Pool) {
	conn := driver.ConnectToPostgresDB()

	return &DBRepository{
		UserRepository: postgresrepo.NewPostgresUserRepo(conn),
	}, conn
}
