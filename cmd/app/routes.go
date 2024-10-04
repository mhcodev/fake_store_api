package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/container"
)

func setupRoutes(app *fiber.App, ch *container.ContainerHandler) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	// ============= User routes ================
	v1.Route("/user", func(router fiber.Router) {
		router.Get("/", ch.UserHandler.GetUsersByParams)
		router.Get("/:id", ch.UserHandler.GetUserByID)
		router.Post("/email/is-available", ch.UserHandler.UserEmailIsAvailable)
		router.Post("/", ch.UserHandler.CreateUser)
		router.Put("/:id", ch.UserHandler.UpdateUser)
		router.Delete("/:id", ch.UserHandler.DeleteUser)
	})

	// ============= Category routes ================
	v1.Route("/category", func(router fiber.Router) {
		router.Get("/", ch.CategoryHandler.GetCategories)
		router.Get("/:id", ch.CategoryHandler.GetCategoryByID)
		router.Post("/", ch.CategoryHandler.CreateCategory)
		router.Put("/:id", ch.CategoryHandler.UpdateCategory)
		router.Delete("/:id", ch.CategoryHandler.DeleteCategory)
	})

	// ============= Category routes ================
	v1.Route("/product", func(router fiber.Router) {
		router.Get("/", ch.ProductHandler.GetProductsByParams)
	})
}
