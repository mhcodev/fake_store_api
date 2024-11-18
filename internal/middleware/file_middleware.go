package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/util"
	"github.com/mhcodev/fake_store_api/internal/validators"
)

func FileSizeLimit(maxFileSize int64) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse form data
		form, err := c.MultipartForm()

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse form data")
		}

		files := form.File["images"]
		errors := validators.ValidationErrors{}

		if len(files) == 0 {
			errors.AddError("files", "Error parsing files")
			return util.ErrorReponse(c, fiber.StatusBadRequest, nil, errors)
		}

		// Validate each file's size
		for _, file := range files {
			if file.Size > maxFileSize {
				maxFileSizeText := maxFileSize / (1024 * 1024)
				msg := fmt.Sprintf("File %s exceeds the allowed size limit: %dMB", file.Filename, maxFileSizeText)
				errors.AddError("files", msg)
				return util.ErrorReponse(c, fiber.StatusBadRequest, nil, errors)
			}
		}

		return c.Next()
	}
}
