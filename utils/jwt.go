package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing JWT (Keep it secret and secure!)
var jwtSecret = []byte("your_secret_key")

// GenerateJWT creates a new JWT token
func GenerateJWT(email string) (string, error) {
	// Define token claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	return token.SignedString(jwtSecret)
}

// VerifyJWT verifies and extracts claims from a JWT token
func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure token is signed with the correct method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
