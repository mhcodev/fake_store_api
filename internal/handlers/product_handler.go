package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
	"github.com/mhcodev/fake_store_api/internal/validators"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

// Get products by params godoc
// @Summary Fetch products by Params
// @Description Fetch products by Params
// @Tags Product
// @Accept json
// @Produce json
// @Param limit query int false "Number of products to return (default 15)"
// @Param offset query int false "Offset for pagination (default 0)"
// @Success 200 {object} models.JSONReponseMany
// @Router /product [get]
func (h *ProductHandler) GetProductsByParams(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 15)
	offset := c.QueryInt("offset", 0)

	parmas := models.QueryParams{
		Limit:  limit,
		Offset: offset,
	}

	products, err := h.ProductService.GetProductsByParams(c.Context(), parmas)
	validationErrors := make(validators.ValidationErrors, 0)

	if err != nil {
		validationErrors.AddError("msg", "no products found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	count, _ := h.ProductService.GetTotalOfProducts(c.Context())

	response := models.JSONReponseMany{
		Success: true,
		Code:    fiber.StatusOK,
		Limit:   limit,
		Offset:  offset,
		Count:   count,
		Data:    products,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Get by ID godoc
// @Summary Get product by ID
// @Description Get product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.JSONReponseOne
// @Router /product/{id} [get]
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id", 0)
	var validationErrors = validators.ValidationErrors{}

	if err != nil {
		validationErrors.AddError("msg", "id is not valid")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	product, err := h.ProductService.GetProductByID(c.Context(), ID)

	if err != nil {
		validationErrors.AddError("msg", "product not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	response := models.JSONReponseOne{
		Success: true,
		Code:    fiber.StatusOK,
		Data:    product,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Create a product godoc
// @Summary Create a product
// @Description Create a product
// @Tags Product
// @Accept json
// @Produce json
// @Param body body services.ProductCreateInput true "Product body request"
// @Success 200 {object} models.JSONReponseOne
// @Router /product [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var input services.ProductCreateInput
	var validationErrors = validators.ValidationErrors{}

	if err := c.BodyParser(&input); err != nil {
		validationErrors.AddError("msg", "error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	validationErrors = validators.ValidateProductCreateInput(input)

	if validationErrors.HasErrors() {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	product, err := h.ProductService.CreateProduct(c.Context(), input)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	response := models.JSONReponseOne{
		Success: true,
		Code:    fiber.StatusOK,
		Data:    product,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Update a product godoc
// @Summary Update a product
// @Description Update a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param body body services.ProductUpdateInput true "Product body request"
// @Success 200 {object} models.JSONReponseOne
// @Router /product/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id", 0)

	var validationErrors = validators.ValidationErrors{}

	if err != nil || ID <= 0 {
		validationErrors.AddError("msg", "id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	var input services.ProductUpdateInput
	if err := c.BodyParser(&input); err != nil {
		validationErrors.AddError("msg", "error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	validationErrors = validators.ValidateProductUpdateInput(input)

	if validationErrors.HasErrors() {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	product, err := h.ProductService.UpdateProduct(c.Context(), ID, input)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	response := models.JSONReponseOne{
		Success: true,
		Code:    fiber.StatusOK,
		Data:    product,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Delete a product godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.JSONReponseOne
// @Router /product/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id", 0)
	var validationErrors = validators.ValidationErrors{}

	if err != nil || ID <= 0 {
		validationErrors.AddError("msg", "id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, validationErrors)
	}

	err = h.ProductService.DeleteProduct(c.Context(), ID)

	if err != nil {
		validationErrors.AddError("msg", err.Error())
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	msg := fmt.Sprintf("product with ID %d deleted", ID)

	response := models.JSONReponseOne{
		Success: true,
		Code:    fiber.StatusOK,
		Data:    msg,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
