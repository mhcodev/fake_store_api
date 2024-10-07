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

type CreateProductRequest struct {
	Product models.Product `json:"product"`
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var request CreateProductRequest
	messages := make([]string, 0)

	if err := c.BodyParser(&request); err != nil {
		messages = append(messages, "error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	product := request.Product
	messages = product.Validate()

	if len(messages) > 0 {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	err := h.ProductService.CreateProduct(c.Context(), &product)

	if err != nil {
		messages = append(messages, "product was not created")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	response := make(map[string]interface{})
	response["product"] = product

	return util.SuccessReponse(c, response)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id", 0)
	messages := make([]string, 0)

	if err != nil || ID <= 0 {
		messages = append(messages, "id is not valid")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	if err != nil {
		messages = append(messages, "product not found")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	var input services.ProductUpdateInput
	if err := c.BodyParser(&input); err != nil {
		messages = append(messages, "error processing request")
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	fmt.Println("input", input)

	validationErrors := validators.ValidateProductUpdateInput(input)

	if validationErrors.HasErrors() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": validationErrors,
		})
	}

	if len(messages) > 0 {
		return util.ErrorReponse(c, fiber.StatusBadRequest, nil, messages)
	}

	product, err := h.ProductService.UpdateProduct(c.Context(), ID, input)

	if err != nil {
		messages = append(messages, "product was not created")
		return util.ErrorReponse(c, fiber.StatusNotFound, nil, messages)
	}

	response := make(map[string]interface{})
	response["product"] = product

	return util.SuccessReponse(c, response)
}
