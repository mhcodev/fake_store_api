package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GeoInfo struct {
	Country string `json:"country_name"`
}

func getCountryFromIP(ip string) (string, error) {
	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var geoInfo GeoInfo
	if err := json.NewDecoder(resp.Body).Decode(&geoInfo); err != nil {
		return "", err
	}

	return geoInfo.Country, nil
}

func getClientIP(c *fiber.Ctx) string {
	ip := c.IP()
	forwarded := c.Get("X-Forwarded-For")
	if forwarded != "" {
		ip = forwarded // Use the forwarded IP if available
	}
	return ip
}
