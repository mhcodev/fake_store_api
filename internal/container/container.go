package container

import (
	"github.com/mhcodev/fake_store_api/internal/handlers"
	"github.com/mhcodev/fake_store_api/internal/repository"
	"github.com/mhcodev/fake_store_api/internal/services"
)

type ContainerService struct {
	// Services
	UserService *services.UserService
}

func NewContainerService(DBRepo *repository.DBRepository) *ContainerService {
	// Initialize Services
	userService := services.NewUserService(&DBRepo.UserRepository)

	// Return the container with all initialized dependencies
	return &ContainerService{
		UserService: userService,
	}
}

type ContainerHandler struct {
	// Handlers
	UserHandler *handlers.UserHandler
}

func NewContainerHandler(cs *ContainerService) *ContainerHandler {
	// Initialize Handlers
	userHandler := handlers.NewUserHandler(cs.UserService)

	// Return the container with all initialized dependencies
	return &ContainerHandler{
		UserHandler: userHandler,
	}
}
