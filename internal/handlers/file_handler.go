package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/services"
	"github.com/mhcodev/fake_store_api/internal/util"
	"github.com/mhcodev/fake_store_api/pkg"
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

	// Array to store uploaded file paths
	var uploadedFiles []string

	// Process each file
	for _, file := range files {
		// Generate a unique filename
		ext := filepath.Ext(file.Filename)
		uniqueName := fmt.Sprintf("%s.%s", pkg.GenerateRandomString(6), ext)

		if ext != ".jpg" && ext != ".png" {
			msg := fmt.Sprintf("Invalid file type for: %s", file.Filename)
			return c.Status(fiber.StatusBadRequest).SendString(msg)
		}

		// Define the upload path
		filePath := filepath.Join("./uploads", uniqueName)

		// Save the file to the server
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file: " + file.Filename)
		}

		fileURL := filepath.Join(c.Hostname(), "./uploads", uniqueName)

		// Add the file path to the response array
		uploadedFiles = append(uploadedFiles, fileURL)
	}

	response := make(map[string]interface{}, 0)
	response["msg"] = "Files uploaded successfully"
	response["files"] = uploadedFiles

	return util.SuccessReponse(c, response)
}
