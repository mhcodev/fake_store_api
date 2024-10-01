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

func (cs *CategoryService) GetCategoryByID(ctx context.Context, ID int) (models.Category, error) {
	return cs.categoryRepository.GetCategoryByID(ctx, ID)
}

func (cs *CategoryService) CreateCategory(ctx context.Context, category *models.Category) error {
	return cs.categoryRepository.CreateCategory(ctx, category)
}

func (cs *CategoryService) UpdateCategory(ctx context.Context, category *models.Category) error {
	return cs.categoryRepository.UpdateCategory(ctx, category)
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, ID int) error {
	return cs.categoryRepository.DeleteCategory(ctx, ID)
}
