package container

import (
	"github.com/mhcodev/fake_store_api/internal/handlers"
	"github.com/mhcodev/fake_store_api/internal/repository"
	"github.com/mhcodev/fake_store_api/internal/services"
)

type ContainerService struct {
	// Services
	UserService     *services.UserService
	CategoryService *services.CategoryService
	ProductService  *services.ProductService
}

func NewContainerService(DBRepo *repository.DBRepository) *ContainerService {
	// Initialize Services
	userService := services.NewUserService(&DBRepo.UserRepository)
	categoryService := services.NewCategoryService(&DBRepo.CategoryRepository)
	productService := services.NewProductService(&DBRepo.ProductRepository)

	// Return the container with all initialized dependencies
	return &ContainerService{
		UserService:     userService,
		CategoryService: categoryService,
		ProductService:  productService,
	}
}

type ContainerHandler struct {
	// Handlers
	UserHandler     *handlers.UserHandler
	CategoryHandler *handlers.CategoryHandler
	ProductHandler  *handlers.ProductHandler
}

func NewContainerHandler(cs *ContainerService) *ContainerHandler {
	// Initialize Handlers
	userHandler := handlers.NewUserHandler(cs.UserService)
	categoryHandler := handlers.NewCategoryHandler(cs.CategoryService)
	productHandler := handlers.NewProductHandler(cs.ProductService)

	// Return the container with all initialized dependencies
	return &ContainerHandler{
		UserHandler:     userHandler,
		CategoryHandler: categoryHandler,
		ProductHandler:  productHandler,
	}
}
