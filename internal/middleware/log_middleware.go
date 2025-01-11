package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/models"
)

func RecordApiLogs(c *fiber.Ctx) error {
	// Process the request
	next := c.Next()

	if strings.Contains(c.Path(), "/metric") {
		return next
	}

	// Record the start time
	startTime := time.Now()

	// Calculate the response time
	duration := time.Since(startTime)

	go func(duration time.Duration, c *fiber.Ctx) {
		ip, err := getClientIP(c)

		if err != nil {
			return
		}

		// Get request details
		country, geoErr := getCountryFromIP(ip)
		if geoErr != nil {
			country = "unknown"
		}

		if strings.TrimSpace(country) == "" {
			country = "unknown"
		}

		l := models.ApiLog{
			Method:       c.Method(),
			Path:         c.Path(),
			Version:      "v1",
			ResponseTime: duration.Milliseconds(),
			UserID:       0,
			IPAdress:     ip,
			Country:      country,
			StatusCode:   200,
		}

		LogService.InsertApiLog(context.Background(), &l)
	}(duration, c)

	return next
}
