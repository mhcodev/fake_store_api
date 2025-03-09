package handlers

import (
	"context"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/services"
)

type FileHandler struct {
	fileService *services.FileService
}

func NewFileHandler(fileService *services.FileService) *FileHandler {
	return &FileHandler{fileService: fileService}
}

// Upload a file godoc
// @Summary Upload a file
// @Description Upload a file
// @Tags File
// @Accept json
// @Produce json
// @Param images formData file true "File to upload"
// @Success 200 {object} models.JSONReponseOne
// @Router /file/upload [post]
func (fh *FileHandler) UploadLoad(c *fiber.Ctx) error {
	// Retrieve the file from the request
	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse form data")
	}

	// Get the array of files with the key "images"
	files := form.File["images"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("No files uploaded")
	}

	// Global context with timeout for the entire operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	results, errors := fh.fileService.ProcessFiles(ctx, files, &wg, c.BaseURL())

	data := make(map[string]any, 0)
	data["msg"] = "Files uploaded successfully"
	data["files"] = results

	if len(errors) > 0 {
		data["errors"] = errors
	}

	response := models.JSONReponseOne{
		Success: true,
		Code:    fiber.StatusOK,
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
