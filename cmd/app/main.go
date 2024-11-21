package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mhcodev/fake_store_api/internal/container"
	"github.com/mhcodev/fake_store_api/internal/middleware"
	"github.com/mhcodev/fake_store_api/internal/repository"
)

const (
	RequestMaxSize = 20 * 1024 * 1024
	UploadDir      = "./uploads"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbType := "postgres"

	var dbRepo *repository.DBRepository

	switch dbType {
	case "postgres":
		postgresRepo, conn := repository.InitPosgresRepositories()
		dbRepo = postgresRepo
		defer conn.Close()
	}

	app := fiber.New(fiber.Config{
		BodyLimit:    RequestMaxSize,
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(middleware.RequestSizeLimit(RequestMaxSize))

	// Ensure the uploads directory exists
	os.MkdirAll(UploadDir, os.ModePerm)
	app.Static("/uploads", UploadDir)

	containerService := container.NewContainerService(dbRepo)
	ch := container.NewContainerHandler(containerService)

	setupRoutes(app, ch)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
