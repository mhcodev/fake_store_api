package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

func (h *OrderHandler) GetOrdersByParams(c *fiber.Ctx) error {
	params := models.QueryParams{
		Limit:  10,
		Offset: 10,
	}

	orders, err := h.OrderService.GetOrdersByParams(c.Context(), params)

	if err != nil {
		return err
	}

	response := make(map[string]interface{}, 0)
	response["orders"] = orders

	return util.SuccessReponse(c, response)
}
