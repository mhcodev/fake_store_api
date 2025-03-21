package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

// @title Fake Store API
// @version 1.0
// @description This is a sample server for a fake store API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /api/v1
func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start the Redis server.
	container.StartRedisServer()

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
		BodyLimit:    RequestMaxSize,
		ErrorHandler: middleware.ErrorHandler,
	})

	middleware.RegisterPrometheusMetrics()

	app.Use(middleware.RequestSizeLimit(RequestMaxSize))
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	app.Use(logger.New())

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

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Not Found",
			"message": "The requested URL " + c.OriginalURL() + " was not found.",
		})
	})

	// Start the server
	log.Fatal(app.Listen(":4000"))
}
