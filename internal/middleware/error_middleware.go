package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Check if the error is related to BodyLimit
	if errors.Is(err, fiber.ErrRequestEntityTooLarge) {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"error": "Request body too large. Maximum allowed size is 20MB.",
		})
	}

	if errors.Is(err, fiber.ErrBadRequest) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})

	}

	if errors.Is(err, fiber.ErrNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Resource not found",
		})
	}

	if errors.Is(err, fiber.ErrForbidden) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}

	if errors.Is(err, fiber.ErrUnauthorized) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	if errors.Is(err, fiber.ErrTooManyRequests) {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Too many requests",
		})
	}

	if errors.Is(err, fiber.ErrServiceUnavailable) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Service Unavailable",
		})
	}

	if errors.Is(err, fiber.ErrInternalServerError) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})

	}

	return c.Next()
}
