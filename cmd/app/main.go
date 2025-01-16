package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/mhcodev/fake_store_api/internal/config"
	"github.com/mhcodev/fake_store_api/internal/container"
	"github.com/mhcodev/fake_store_api/internal/driver"
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
	var conn *pgxpool.Pool

	switch dbType {
	case "postgres":
		conn := driver.ConnectToPostgresDB()
		dbRepo = repository.InitPosgresRepositories(conn)
	}

	defer conn.Close()

	app := fiber.New(fiber.Config{
		BodyLimit: RequestMaxSize,
	})

	middleware.RegisterPrometheusMetrics()

	app.Use(middleware.RequestSizeLimit(RequestMaxSize))

	// Ensure the uploads directory exists
	os.MkdirAll(UploadDir, os.ModePerm)
	app.Static("/uploads", UploadDir)

	containerService := container.NewContainerService(dbRepo)
	ch := container.NewContainerHandler(containerService)

	middleware.LogService = containerService.LogService

	app.Use(middleware.RecordApiLogs)

	// Save connection variable in app configuration
	config.NewAppConfiguration(conn)

	setupRoutes(app, ch)
	registerPrometheusRoute(app)

	app.Use(middleware.ErrorHandler)

	// Start the server
	log.Fatal(app.Listen(":4000"))
}
