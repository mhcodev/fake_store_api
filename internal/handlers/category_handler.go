package handlers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mhcodev/fake_store_api/docs"
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

// Get Categories godoc
// @Summary Fetch all categories
// @Description Fetch all categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} models.JSONReponseMany
// @Router /category [get]
func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.CategoryService.GetCategories(c.Context())
	var validationErrors = validators.ValidationErrors{}

	if err != nil {
		validationErrors.AddError("errors", err.Error())
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	count, _ := h.CategoryService.GetTotalOfCategories(c.Context())

	response := models.JSONReponseMany{
		Success: true,
		Count:   count,
		Data:    categories,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Get Category by ID godoc
// @Summary Fetch category by ID
// @Description Fetch category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} models.JSONReponseOne
// @Router /category [get]
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

	response := models.JSONReponseOne{
		Success: true,
		Data:    category,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

type CreateCategoryRequest struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
}

// Create Category godoc
// @Summary Create a category
// @Description Create a category
// @Tags Category
// @Accept json
// @Produce json
// @Param body body CreateCategoryRequest true "Category body request"
// @Success 200 {object} models.JSONReponseOne
// @Router /category [post]
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var request CreateCategoryRequest
	var validationErrors = validators.ValidationErrors{}

	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		validationErrors.AddError("errors", "Error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	if strings.TrimSpace(request.Name) == "" {
		validationErrors.AddError("name", "name is required")
	}

	if strings.TrimSpace(request.ImageURL) == "" {
		validationErrors.AddError("imageURL", "imageURL is required")
	}

	validImage, _ := pkg.IsImageURL(strings.TrimSpace(request.ImageURL))

	if !validImage {
		validationErrors.AddError("imageURL", "imageURL has to contains a valid image")
	}

	if validationErrors.HasErrors() {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	category := models.Category{
		Name:     request.Name,
		ImageURL: request.ImageURL,
	}

	err := h.CategoryService.CreateCategory(c.Context(), &category)

	if err != nil {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	response := models.JSONReponseOne{
		Success: true,
		Data:    category,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

type UpdateCategoryRequest struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
}

// Update Category godoc
// @Summary Update a category
// @Description Update a category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param body body UpdateCategoryRequest true "Category body request"
// @Success 200 {object} models.JSONReponseOne
// @Router /category/{id} [put]
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

	if strings.TrimSpace(request.Name) != "" {
		category.Name = request.Name
	}

	if strings.TrimSpace(request.ImageURL) != "" {
		category.ImageURL = request.ImageURL

		validImage, _ := pkg.IsImageURL(strings.TrimSpace(request.ImageURL))

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

// Delete Categor by ID godoc
// @Summary Delete a category
// @Description Delete a category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.JSONReponseOne
// @Router /category/{id} [delete]
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
