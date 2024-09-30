package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
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

type CreateCategoryRequest struct {
	Category models.Category `json:"category"`
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var request CreateCategoryRequest
	messages := make([]string, 0)

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		messages = append(messages, "Error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	category := request.Category

	if strings.TrimSpace(category.Name) == "" {
		messages = append(messages, "name is required")
	}

	if strings.TrimSpace(category.ImageURL) == "" {
		messages = append(messages, "image_url is required")
	}

	validImage, _ := util.IsImageURL(strings.TrimSpace(category.ImageURL))

	if !validImage {
		messages = append(messages, "image_url has to contains a valid image")
	}

	if len(messages) > 0 {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	err := h.CategoryService.CreateCategory(c.Context(), &category)

	if err != nil {
		messages = append(messages, err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	response := fiber.Map{"category": category}

	return util.SuccessReponse(c, response)
}
