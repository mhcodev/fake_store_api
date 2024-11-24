package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
	"github.com/mhcodev/fake_store_api/internal/repository/repositories"
	"github.com/mhcodev/fake_store_api/pkg"
)

const (
	MaxConcurrentUploads = 5
	UploadDir            = "./uploads"
	MaxFileSize          = 5 * 1024 * 1024
)

type FileService struct {
	semaphore      chan struct{}
	fileRepository repositories.FileRepository
}

func NewFileService(fileRepository *repositories.FileRepository) *FileService {
	return &FileService{
		fileRepository: *fileRepository,
		semaphore:      make(chan struct{}, MaxConcurrentUploads),
	}
}

type SaveFileInput struct {
	ctx     context.Context
	c       *fiber.Ctx
	File    *multipart.FileHeader
	BaseURL string
	wg      *sync.WaitGroup
}

type ProcessFilesInput struct {
	Files   []*multipart.FileHeader
	BaseURL string
}

func (fs *FileService) SaveFileToDB(ctx context.Context, file *models.File) error {
	baseURL, err := pkg.GetBaseURL(file.Url)
	file.BaseURL = baseURL

	if err != nil {
		return nil
	}

	return fs.fileRepository.SaveFileToDB(ctx, file)
}

func (fs *FileService) SaveFileToDisk(ctx context.Context, file *multipart.FileHeader, baseURL string) (models.File, error) {
	// Generate unique file name
	ext := filepath.Ext(file.Filename)
	uniqueName := fmt.Sprintf("%s%s", pkg.GenerateRandomString(6), ext)
	filePath := filepath.Join(UploadDir, uniqueName)

	// Create a new context with a timeout for the file-saving process
	saveCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan error, 1)

	go func() {
		// Open the file
		src, err := file.Open()
		if err != nil {
			done <- fmt.Errorf("failed to open file %s: %v", file.Filename, err)
			return
		}
		defer src.Close()

		// Create destination file
		dest, err := os.Create(filePath)
		if err != nil {
			done <- fmt.Errorf("failed to create file %s: %v", filePath, err)
			return
		}
		defer dest.Close()

		// Copy file content
		if _, err := dest.ReadFrom(src); err != nil {
			done <- fmt.Errorf("failed to copy content to file %s: %v", filePath, err)
			return
		}

		done <- nil
	}()

	// Generate public URL
	publicURL := fmt.Sprintf("%s/uploads/%s", baseURL, uniqueName)

	select {
	case <-saveCtx.Done():
		msg := fmt.Sprintf("File saving timed out for %s\n", file.Filename)
		log.Println(msg)
		return models.File{}, errors.New(msg)
	case err := <-done:
		if err != nil {
			msg := fmt.Sprintf("Error saving file: %v\n", err)
			log.Println(msg)
			return models.File{}, errors.New(msg)
		}
	}

	return models.File{
		OriginalName: file.Filename,
		FileName:     uniqueName,
		Type:         ext,
		Url:          publicURL,
	}, nil
}

func (fs *FileService) ProcessFiles(ctx context.Context, files []*multipart.FileHeader, wg *sync.WaitGroup, baseURL string) ([]models.File, []string) {
	resultCh := make(chan models.File, len(files))
	errorCh := make(chan error, len(files))

	fmt.Println("====================================")
	fmt.Println("FILES")
	fmt.Println("====================================")
	for _, f := range files {
		fmt.Println(f.Filename)

		wg.Add(1)
		go func() {
			defer wg.Done()
			newFile, err := fs.SaveFileToDisk(ctx, f, baseURL)

			if err != nil {
				errorCh <- err
			} else {
				resultCh <- newFile
			}
		}()
	}
	fmt.Println("====================================")

	wg.Add(1)
	results := make([]models.File, 0)
	errors := make([]string, 0)
	var totalProccesed int
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Database saving canceled or timed out")
				return
			case file, ok := <-resultCh:
				if !ok {
					// Channel closed, exit the loop
					return
				}

				err := fs.SaveFileToDB(ctx, &file)

				if err != nil {
					fmt.Println(err.Error())
					errors = append(errors, err.Error())
				} else {
					results = append(results, file)
					fmt.Printf("(%d) %s saved in db \n", totalProccesed, file.OriginalName)
				}

				totalProccesed++

				if totalProccesed == len(files) {
					return
				}

			}
		}
	}()

	wg.Wait()

	close(resultCh)
	close(errorCh)

	if len(errors) > 0 {
		return results, errors
	}

	return results, []string{}
}

// func (fs *FileService) GetFileByName(ctx context.Context, filename string) (models.File, error) {
// 	return fs.fileRepository.GetFileByName(filename)
// }
