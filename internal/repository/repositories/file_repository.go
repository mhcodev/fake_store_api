package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type FileRepository interface {
	UploadFile(ctx context.Context, file *models.File) error
	// GetFileByName(filename string) (models.File, error)
	// DeleteFileByName(filename string) error
}
