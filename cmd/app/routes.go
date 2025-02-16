package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/mhcodev/fake_store_api/internal/container"
	"github.com/mhcodev/fake_store_api/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func setupRoutes(app *fiber.App, ch *container.ContainerHandler) {
	api := app.Group("/api",
		middleware.RecordRequestLatency,
		middleware.RecordRequestCount,
		middleware.RecordRequestFrequency,
	)

	v1 := api.Group("/v1")

	// ============= Auth routes ================
	v1.Route("/auth", func(router fiber.Router) {
		router.Post("/login", ch.AuthHandler.Login)
		router.Get("/data", ch.AuthHandler.GetTokenData)
		router.Post("/refresh", ch.AuthHandler.AccessTokenFromRefreshToken)
	})

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

	// ============= Product routes ================
	v1.Route("/product", func(router fiber.Router) {
		router.Get("/", ch.ProductHandler.GetProductsByParams)
		router.Get("/:id", ch.ProductHandler.GetProductByID)
		router.Post("/", ch.ProductHandler.CreateProduct)
		router.Put("/:id", ch.ProductHandler.UpdateProduct)
		router.Delete("/:id", ch.ProductHandler.DeleteProduct)
	})

	// ============= Auth routes ================
	v1.Route("/file", func(router fiber.Router) {
		router.Post("/upload",
			middleware.FileSizeLimit(5*1024*1024),
			ch.FileHandler.UploadLoad,
		)
	})
}

func registerPrometheusRoute(app *fiber.App) {
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
}
