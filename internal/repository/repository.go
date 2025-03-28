package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	postgresrepository "github.com/mhcodev/fake_store_api/internal/repository/postgres_repository"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type DBRepository struct {
	AuthRepository     repositories.AuthRepository
	UserRepository     repositories.UserRepository
	CategoryRepository repositories.CategoryRepository
	ProductRepository  repositories.ProductRepository
	FileRepository     repositories.FileRepository
	LogRepository      repositories.LogRepository
}

func InitPosgresRepositories(conn *pgxpool.Pool) *DBRepository {
	return &DBRepository{
		AuthRepository:     postgresrepository.NewPostgresAuthRepository(conn),
		UserRepository:     postgresrepository.NewPostgresUserRepository(conn),
		CategoryRepository: postgresrepository.NewPostgresCategoryRepository(conn),
		ProductRepository:  postgresrepository.NewPostgresProductRepository(conn),
		FileRepository:     postgresrepository.NewPostgresFileRepository(conn),
		LogRepository:      postgresrepository.NewPostgresLogRepository(conn),
	}
}
