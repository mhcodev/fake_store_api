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

func (h *ProductHandler) GetProductsByParams(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 15)
	offset := c.QueryInt("offset", 0)

	parmas := models.QueryParams{
		Limit:  limit,
		Offset: offset,
	}

	products, err := h.ProductService.GetProductsByParams(c.Context(), parmas)
	var validationErrors validators.ValidationErrors

	if err != nil {
		validationErrors.AddError("msg", "no products found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	count, _ := h.ProductService.GetTotalOfProducts(c.Context())

	response := make(map[string]interface{})
	response["count"] = count
	response["products"] = products

	return util.SuccessReponse(c, response)
}

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

	response := make(map[string]interface{})
	response["product"] = product

	return util.SuccessReponse(c, response)
}

type CreateProductRequest struct {
	Product models.Product `json:"product"`
}

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
		validationErrors.AddError("msg", "product was not created")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, validationErrors)
	}

	response := make(map[string]interface{})
	response["product"] = product

	return util.SuccessReponse(c, response)
}

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

	response := make(map[string]interface{})
	response["product"] = product

	return util.SuccessReponse(c, response)
}

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

	response := make(map[string]interface{})
	response["msg"] = fmt.Sprintf("product %d deleted", ID)

	return util.SuccessReponse(c, response)
}
