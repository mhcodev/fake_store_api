package handlers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
	"github.com/mhcodev/fake_store_api/internal/validators"
	"github.com/mhcodev/fake_store_api/pkg"
)

type CategoryHandler struct {
	CategoryService *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: categoryService}
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.CategoryService.GetCategories(c.Context())
	var validationErrors = validators.ValidationErrors{}

	if err != nil {
		validationErrors.AddError("errors", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	count, _ := h.CategoryService.GetTotalOfCategories(c.Context())

	response := fiber.Map{
		"count":      count,
		"categories": categories,
	}

	return util.SuccessReponse(c, response)
}

func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("id", "0"))
	var validationErrors = validators.ValidationErrors{}

	if err != nil || categoryID == 0 {
		validationErrors.AddError("errors", "category id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	category, err := h.CategoryService.GetCategoryByID(c.Context(), categoryID)

	if err != nil {
		validationErrors.AddError("errors", "category not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	response := fiber.Map{"category": category}

	return util.SuccessReponse(c, response)
}

type CreateCategoryRequest struct {
	Category models.Category `json:"category"`
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var request CreateCategoryRequest
	var validationErrors = validators.ValidationErrors{}

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		validationErrors.AddError("errors", "Error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	category := request.Category

	if strings.TrimSpace(category.Name) == "" {
		validationErrors.AddError("name", "name is required")
	}

	if strings.TrimSpace(category.ImageURL) == "" {
		validationErrors.AddError("imageURL", "imageURL is required")
	}

	validImage, _ := pkg.IsImageURL(strings.TrimSpace(category.ImageURL))

	if !validImage {
		validationErrors.AddError("imageURL", "imageURL has to contains a valid image")
	}

	if validationErrors.HasErrors() {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	err := h.CategoryService.CreateCategory(c.Context(), &category)

	if err != nil {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := fiber.Map{"category": category}

	return util.SuccessReponse(c, response)
}

type UpdateCategoryRequest struct {
	Category models.Category `json:"category"`
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	var request UpdateCategoryRequest
	var validationErrors = validators.ValidationErrors{}

	categoryID, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil || categoryID == 0 {
		validationErrors.AddError("category", "category id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	category, err := h.CategoryService.GetCategoryByID(c.Context(), categoryID)

	if err != nil {
		validationErrors.AddError("category", "category not found")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		validationErrors.AddError("category", "error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	if strings.TrimSpace(request.Category.Name) != "" {
		category.Name = request.Category.Name
	}

	if strings.TrimSpace(request.Category.ImageURL) != "" {
		category.ImageURL = request.Category.ImageURL

		validImage, _ := pkg.IsImageURL(strings.TrimSpace(request.Category.ImageURL))

		if !validImage {
			validationErrors.AddError("imageURL", "imageURL has to contains a valid image")
		}
	}

	if validationErrors.HasErrors() {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	err = h.CategoryService.UpdateCategory(c.Context(), &category)

	if err != nil {
		validationErrors.AddError("category", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := fiber.Map{"category": category}

	return util.SuccessReponse(c, response)
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	var validationErrors = validators.ValidationErrors{}

	categoryID, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil || categoryID == 0 {
		validationErrors.AddError("category", "category id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	_, err = h.CategoryService.GetCategoryByID(c.Context(), categoryID)

	if err != nil {
		validationErrors.AddError("category", "category not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	err = h.CategoryService.DeleteCategory(c.Context(), categoryID)

	if err != nil {
		validationErrors.AddError("category", "category was not deleted")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	if validationErrors.HasErrors() {
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	response := fiber.Map{"msg": "category deleted"}

	return util.SuccessReponse(c, response)
}
