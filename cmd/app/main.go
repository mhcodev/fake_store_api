package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/mhcodev/fake_store_api/internal/container"
	"github.com/mhcodev/fake_store_api/internal/middleware"
	"github.com/mhcodev/fake_store_api/internal/repository"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RequestMaxSize = 20 * 1024 * 1024
	UploadDir      = "./uploads"
)

// Define metrics
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_request_count",
			Help: "Total number of requests",
		},
		[]string{"method", "endpoint"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Histogram of response time for API requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}

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
		dbRepo, conn = repository.InitPosgresRepositories()
	}

	defer conn.Close()

	app := fiber.New(fiber.Config{
		BodyLimit:    RequestMaxSize,
		ErrorHandler: middleware.ErrorHandler,
	})

	middleware.RegisterPrometheusMetrics()

	app.Use(middleware.RequestSizeLimit(RequestMaxSize))

	// Ensure the uploads directory exists
	os.MkdirAll(UploadDir, os.ModePerm)
	app.Static("/uploads", UploadDir)

	containerService := container.NewContainerService(dbRepo)
	ch := container.NewContainerHandler(containerService)

	setupRoutes(app, ch)
	registerPrometheusRoute(app)

	// Start the server
	log.Fatal(app.Listen(":4000"))
}
