package middleware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mhcodev/fake_store_api/internal/container"
	"github.com/mhcodev/fake_store_api/internal/models"
)

func RecordApiLogs(c *fiber.Ctx) error {
	// Record the start time
	startTime := time.Now()

	// Process the request
	next := c.Next()

	if strings.Contains(c.Path(), "/metric") {
		return next
	}

	// Calculate the response time
	duration := time.Since(startTime)

	method := c.Method()
	path := c.Path()
	statusCode := c.Response().StatusCode()

	go func(duration time.Duration, method, path string, statusCode int) {
		ip, err := getClientIP()

		if err != nil {
			return
		}

		countryStored := ""
		ctx := context.Background()
		key := fmt.Sprintf("geo:info:%s", ip)
		redis, redisErr := container.GetRedisClient()

		if redisErr != nil {
			fmt.Println(redisErr)
		} else {
			countryMap, err := redis.GetAll(ctx, key)

			if err != nil {
				countryStored = ""
			} else {
				countryStored = countryMap["name"]
			}
		}

		if countryStored == "" {
			// Get request details
			country, geoErr := getCountryFromIP(ip)

			if geoErr != nil {
				country = "unknown"
			}

			if strings.TrimSpace(country) == "" {
				country = "unknown"
			}

			countryStored = country

			if redisErr == nil {
				countryMap := map[string]interface{}{
					"name": countryStored,
				}
				redis.Set(ctx, key, countryMap)
			}
		}

		l := models.ApiLog{
			Method:       method,
			Path:         path,
			Version:      "v1",
			ResponseTime: duration.Milliseconds(),
			UserID:       0,
			IPAdress:     ip,
			Country:      countryStored,
			StatusCode:   statusCode,
		}

		LogService.InsertApiLog(context.Background(), &l)
	}(duration, method, path, statusCode)

	return next
}
