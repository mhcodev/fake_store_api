package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
)

// ErrorReponse return a server error and request status
func ErrorReponse(c *fiber.Ctx, statusCode int, data any, errors map[string][]string) error {
	response := fiber.Map{
		"success": false,
		"code":    statusCode,
		"errors":  errors,
	}

	if data != nil {
		response["data"] = data
	}

	return c.Status(statusCode).JSON(response)
}

// SuccessReponse returns a success response when a request is successful
func SuccessReponseOne(c *fiber.Ctx, data any) error {
	response := models.JSONReponseOne{
		Success: true,
		Code:    fiber.StatusOK,
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// SuccessReponse returns a success response when a request is successful
func SuccessReponse(c *fiber.Ctx, data any) error {
	response := models.JSONReponseOne{
		Success: true,
		Code:    fiber.StatusOK,
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
