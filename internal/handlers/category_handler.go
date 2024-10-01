package handlers

import (
	"strconv"
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

func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	messages := make([]string, 0)
	categoryID, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil || categoryID == 0 {
		messages = append(messages, "category id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	category, err := h.CategoryService.GetCategoryByID(c.Context(), categoryID)

	if err != nil {
		messages = append(messages, "category not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	response := fiber.Map{"category": category}

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
		messages = append(messages, "imageURL is required")
	}

	validImage, _ := util.IsImageURL(strings.TrimSpace(category.ImageURL))

	if !validImage {
		messages = append(messages, "imageURL has to contains a valid image")
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

type UpdateCategoryRequest struct {
	Category models.Category `json:"category"`
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	var request UpdateCategoryRequest
	messages := make([]string, 0)

	categoryID, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil || categoryID == 0 {
		messages = append(messages, "category id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	category, err := h.CategoryService.GetCategoryByID(c.Context(), categoryID)

	if err != nil {
		messages = append(messages, "category not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		messages = append(messages, "Error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	if strings.TrimSpace(request.Category.Name) != "" {
		category.Name = request.Category.Name
	}

	if strings.TrimSpace(request.Category.ImageURL) != "" {
		category.ImageURL = request.Category.ImageURL

		validImage, _ := util.IsImageURL(strings.TrimSpace(request.Category.ImageURL))

		if !validImage {
			messages = append(messages, "imageURL has to contains a valid image")
		}
	}

	if len(messages) > 0 {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	err = h.CategoryService.UpdateCategory(c.Context(), &category)

	if err != nil {
		messages = append(messages, err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	response := fiber.Map{"category": category}

	return util.SuccessReponse(c, response)
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	messages := make([]string, 0)

	categoryID, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil || categoryID == 0 {
		messages = append(messages, "category id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	_, err = h.CategoryService.GetCategoryByID(c.Context(), categoryID)

	if err != nil {
		messages = append(messages, "category not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	err = h.CategoryService.DeleteCategory(c.Context(), categoryID)

	if err != nil {
		messages = append(messages, "category was not deleted")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	if len(messages) > 0 {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	response := fiber.Map{"msg": "category deleted"}

	return util.SuccessReponse(c, response)
}
