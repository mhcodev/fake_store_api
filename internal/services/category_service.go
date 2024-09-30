package services

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type CategoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: *categoryRepository,
	}
}

func (cs *CategoryService) GetCategories(ctx context.Context) ([]models.Category, error) {
	return cs.categoryRepository.GetCategories(ctx)
}
