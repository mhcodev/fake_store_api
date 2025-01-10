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

	// Record the start time
	startTime := time.Now()

	// Calculate the response time
	duration := time.Since(startTime)

	ip := getClientIP(c)

	go func(ipReq string) {
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
	}(ip)

	return next
}
