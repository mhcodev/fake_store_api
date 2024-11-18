package services

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type FileService struct {
	fileRepository repositories.FileRepository
}

func NewFileService(fileRepository *repositories.FileRepository) *FileService {
	return &FileService{
		fileRepository: *fileRepository,
	}
}

func (fs *FileService) UploadFile(ctx context.Context, file *models.File) error {
	return fs.fileRepository.UploadFile(ctx, file)
}

// func (fs *FileService) GetFileByName(ctx context.Context, filename string) (models.File, error) {
// 	return fs.fileRepository.GetFileByName(filename)
// }
