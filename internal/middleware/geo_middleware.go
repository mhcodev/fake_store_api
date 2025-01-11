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

type IpInfo struct {
	IP string `json:"ip"`
}

func getClientIP(c *fiber.Ctx) (string, error) {
	url := fmt.Sprintf("https://api.ipify.org?format=json")
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ipInfo IpInfo
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		return "", err
	}

	return ipInfo.IP, nil
}
