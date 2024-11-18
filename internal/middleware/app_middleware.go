package middleware

import "github.com/gofiber/fiber/v2"

func RequestSizeLimit(maxSize int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check Content-Length header
		contentLength := c.Get("Content-Length")
		if contentLength != "" {
			if c.Request().Header.ContentLength() > maxSize {
				return c.Status(fiber.StatusRequestEntityTooLarge).SendString("Request size exceeds the allowed limit")
			}
		}
		return c.Next()
	}
}
