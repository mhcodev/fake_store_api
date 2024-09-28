package util

import "github.com/gofiber/fiber/v2"

func ErrorReponse(c *fiber.Ctx, statusCode int, data interface{}, messages []string) error {
	response := fiber.Map{
		"success":    false,
		"statusCode": statusCode,
		"error":      fiber.ErrBadRequest.Error(),
		"messages":   messages,
	}

	if data != nil {
		response["data"] = data
	}

	return c.Status(fiber.StatusNotFound).JSON(response)
}

func SuccessReponse(c *fiber.Ctx, data map[string]interface{}) error {
	response := fiber.Map{
		"success":    true,
		"statusCode": 200,
	}

	for key, value := range data {
		response[key] = value
	}

	return c.Status(fiber.StatusNotFound).JSON(response)
}

// Function to check if a number is in the slice
func Includes(arr []int, num int) bool {
	for _, value := range arr {
		if value == num {
			return true
		}
	}
	return false
}
