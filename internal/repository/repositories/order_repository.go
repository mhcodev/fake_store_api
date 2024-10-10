package repositories

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
)

type OrderRepository interface {
	GetOrdersByParams(ctx context.Context, params models.QueryParams) ([]models.Order, error)
}
