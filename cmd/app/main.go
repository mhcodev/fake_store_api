package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mhcodev/fake_store_api/internal/container"
	"github.com/mhcodev/fake_store_api/internal/repository"
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

	app := fiber.New()

	containerService := container.NewContainerService(dbRepo)
	ch := container.NewContainerHandler(containerService)

	fmt.Println(ch)

	setupRoutes(app, ch)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
