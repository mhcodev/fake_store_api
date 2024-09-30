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
}

func NewContainerService(DBRepo *repository.DBRepository) *ContainerService {
	// Initialize Services
	userService := services.NewUserService(&DBRepo.UserRepository)
	categoryService := services.NewCategoryService(&DBRepo.CategoryRepository)

	// Return the container with all initialized dependencies
	return &ContainerService{
		UserService:     userService,
		CategoryService: categoryService,
	}
}

type ContainerHandler struct {
	// Handlers
	UserHandler     *handlers.UserHandler
	CategoryHandler *handlers.CategoryHandler
}

func NewContainerHandler(cs *ContainerService) *ContainerHandler {
	// Initialize Handlers
	userHandler := handlers.NewUserHandler(cs.UserService)
	categoryHandler := handlers.NewCategoryHandler(cs.CategoryService)

	// Return the container with all initialized dependencies
	return &ContainerHandler{
		UserHandler:     userHandler,
		CategoryHandler: categoryHandler,
	}
}
