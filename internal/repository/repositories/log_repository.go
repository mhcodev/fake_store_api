package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type LogRepository interface {
	InsertApiLog(ctx context.Context, apiLog *models.ApiLog) error
}
