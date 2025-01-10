package services

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type LogService struct {
	logRepository repositories.LogRepository
}

func NewLogService(logRepository *repositories.LogRepository) *LogService {
	return &LogService{
		logRepository: *logRepository,
	}
}

func (ls *LogService) InsertApiLog(ctx context.Context, apiLog *models.ApiLog) error {
	return ls.logRepository.InsertApiLog(ctx, apiLog)
}
