package handlers

import (
	"context"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
)

type FileHandler struct {
	fileService *services.FileService
}

func NewFileHandler(fileService *services.FileService) *FileHandler {
	return &FileHandler{fileService: fileService}
}

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

	response := make(map[string]interface{}, 0)
	response["msg"] = "Files uploaded successfully"
	response["files"] = results
	response["errors"] = errors

	return util.SuccessReponse(c, response)
}
