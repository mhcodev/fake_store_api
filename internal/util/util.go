package util

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ErrorReponse return a server error and request status
func ErrorReponse(c *fiber.Ctx, statusCode int, data interface{}, messages []string) error {
	response := fiber.Map{
		"success":  false,
		"code":     statusCode,
		"error":    fiber.ErrBadRequest.Error(),
		"messages": messages,
	}

	if data != nil {
		response["data"] = data
	}

	return c.Status(fiber.StatusNotFound).JSON(response)
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

	return c.Status(fiber.StatusNotFound).JSON(response)
}

// Includes checks if a number is in the slice
func Includes(arr []int, num int) bool {
	for _, value := range arr {
		if value == num {
			return true
		}
	}
	return false
}

// IsImageURL checks if the Content-Type starts with "image/"
func IsImageURL(url string) (bool, error) {
	resp, err := http.Head(url)
	if err != nil {
		return false, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return false, fmt.Errorf("could not find Content-Type header")
	}

	// Check if the Content-Type starts with "image/"
	return strings.HasPrefix(contentType, "image/"), nil
}
