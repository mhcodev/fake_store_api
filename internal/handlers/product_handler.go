package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) GetProductsByParams(c *fiber.Ctx) error {
	products, err := h.ProductService.GetProductsByParams(c.Context(), models.QueryParams{})

	messages := make([]string, 0)

	if err != nil {
		messages = append(messages, "no products found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	response := make(map[string]interface{})
	response["products"] = products

	return util.SuccessReponse(c, response)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id", 0)
	messages := make([]string, 0)

	if err != nil {
		messages = append(messages, "id is not valid")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	product, err := h.ProductService.GetProductByID(c.Context(), ID)

	if err != nil {
		messages = append(messages, "product not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	response := make(map[string]interface{})
	response["product"] = product

	return util.SuccessReponse(c, response)
}
