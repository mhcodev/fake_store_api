package pkg

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("pepeXD") // Replace with a secure secret key

const accessTokenExpiry = time.Minute * 15    // Access token valid for 15 minutes
const refreshTokenExpiry = time.Hour * 24 * 7 // Refresh token valid for 7 days

// GenerateAccessToken generates a JWT access token
func GenerateAccessToken(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(accessTokenExpiry).Unix(),
	}

	for k, v := range data {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret) // Sign with the secret key
}

// GenerateRefreshToken generates a JWT refresh token
func GenerateRefreshToken(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(refreshTokenExpiry).Unix(), // Token expiration
	}

	for k, v := range data {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret) // Sign with the secret key
}

// ValidateToken returns a token validated
func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil // Use the same secret to validate the token
	})
}

// Function to parse JWT and extract claims
func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	// Parse the JWT token
	token, err := ValidateToken(tokenStr)

	if err != nil {
		return nil, err
	}

	// If the token is valid, extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	// If token parsing failed, return the error
	return nil, err
}
