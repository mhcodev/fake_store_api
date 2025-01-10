package container

import (
	"github.com/mhcodev/fake_store_api/internal/handlers"
	"github.com/mhcodev/fake_store_api/internal/repository"
	"github.com/mhcodev/fake_store_api/internal/services"
)

type ContainerService struct {
	// Services
	AuthService     *services.AuthService
	UserService     *services.UserService
	CategoryService *services.CategoryService
	ProductService  *services.ProductService
	FileService     *services.FileService
	LogService      *services.LogService
}

func NewContainerService(DBRepo *repository.DBRepository) *ContainerService {
	// Initialize Services
	authService := services.NewAuthService(&DBRepo.AuthRepository)
	userService := services.NewUserService(&DBRepo.UserRepository)
	categoryService := services.NewCategoryService(&DBRepo.CategoryRepository)
	productService := services.NewProductService(&DBRepo.ProductRepository, &DBRepo.CategoryRepository)
	fileService := services.NewFileService(&DBRepo.FileRepository)
	logService := services.NewLogService(&DBRepo.LogRepository)

	// Return the container with all initialized dependencies
	return &ContainerService{
		AuthService:     authService,
		UserService:     userService,
		CategoryService: categoryService,
		ProductService:  productService,
		FileService:     fileService,
		LogService:      logService,
	}
}

type ContainerHandler struct {
	// Handlers
	AuthHandler     *handlers.AuthHandler
	UserHandler     *handlers.UserHandler
	CategoryHandler *handlers.CategoryHandler
	ProductHandler  *handlers.ProductHandler
	OrderHandler    *handlers.OrderHandler
	FileHandler     *handlers.FileHandler
}

func NewContainerHandler(cs *ContainerService) *ContainerHandler {
	// Initialize Handlers
	authHandler := handlers.NewAuthHandler(cs.AuthService)
	userHandler := handlers.NewUserHandler(cs.UserService)
	categoryHandler := handlers.NewCategoryHandler(cs.CategoryService)
	productHandler := handlers.NewProductHandler(cs.ProductService)
	fileHandler := handlers.NewFileHandler(cs.FileService)

	// Return the container with all initialized dependencies
	return &ContainerHandler{
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
		CategoryHandler: categoryHandler,
		ProductHandler:  productHandler,
		FileHandler:     fileHandler,
	}
}
