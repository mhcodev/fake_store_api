package services

import (
	"context"

	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
)

type OrderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: *orderRepository,
	}
}

func (os *OrderService) GetOrdersByParams(ctx context.Context, params models.QueryParams) ([]models.Order, error) {
	orders, err := os.orderRepository.GetOrdersByParams(ctx, params)

	if err != nil {
		return []models.Order{}, err
	}

	return orders, nil
}
