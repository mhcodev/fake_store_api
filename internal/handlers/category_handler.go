package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
)

type CategoryHandler struct {
	CategoryService *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: categoryService}
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.CategoryService.GetCategories(c.Context())

	messages := make([]string, 0)

	if err != nil {
		messages = append(messages, err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	response := fiber.Map{"categories": categories}

	return util.SuccessReponse(c, response)
}
