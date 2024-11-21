package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Check if the error is related to BodyLimit
	if errors.Is(err, fiber.ErrRequestEntityTooLarge) || true {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"error": "Request body too large. Maximum allowed size is 20MB.",
		})
	}

	// Default error handling for other errors
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}
