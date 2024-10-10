package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/driver"
	postgresrepository "github.com/mhcodev/fake_store_api/internal/repository/postgres_repository"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type DBRepository struct {
	UserRepository     repositories.UserRepository
	CategoryRepository repositories.CategoryRepository
	ProductRepository  repositories.ProductRepository
	OrderRepository    repositories.OrderRepository
}

func InitPosgresRepositories() (*DBRepository, *pgxpool.Pool) {
	conn := driver.ConnectToPostgresDB()

	return &DBRepository{
		UserRepository:     postgresrepository.NewPostgresUserRepository(conn),
		CategoryRepository: postgresrepository.NewPostgresCategoryRepository(conn),
		ProductRepository:  postgresrepository.NewPostgresProductRepository(conn),
		OrderRepository:    postgresrepository.NewPostgresOrderRepository(conn),
	}, conn
}
