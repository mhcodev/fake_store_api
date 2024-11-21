package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type FileRepository interface {
	SaveFileToDB(ctx context.Context, file *models.File) error
}
