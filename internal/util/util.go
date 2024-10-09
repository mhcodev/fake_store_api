package util

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorReponse return a server error and request status
func ErrorReponse(c *fiber.Ctx, statusCode int, data interface{}, errors map[string][]string) error {
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
func SuccessReponse(c *fiber.Ctx, data map[string]interface{}) error {
	response := fiber.Map{
		"success": true,
		"code":    200,
	}

	for key, value := range data {
		response[key] = value
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
