package services

import (
	"context"
	"errors"
	"net/http"
	"personal-dashboard-backend/api/user/model"
	"personal-dashboard-backend/db"
	"personal-dashboard-backend/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool { 
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUserService(c *gin.Context) error { 
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return err
	}

	hashedPassword, passWordErr := HashPassword(user.Password)
	if passWordErr != nil {
		return errors.New("failed to hash password")
	}

	_, err := db.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, hashedPassword)
	if err != nil {
		return errors.New("failed to register user")
	}

	return nil
}

func LoginUserService(c *gin.Context) error {
	var loginReq model.LoginRequest
	var storedPassword string

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		return err
	}

	err := db.DB.QueryRow("SELECT password FROM users WHERE email = $1", loginReq.Email).Scan(&storedPassword)
	if err != nil { 
		return errors.New("user not found")
	}

	if !CheckPasswordHash(loginReq.Password, storedPassword) {
		return errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(loginReq.Email)
	if err != nil {
		return errors.New("failed to generate token")
	}

	resp := model.LoginResponse{Token : token}
	c.JSON(http.StatusOK, resp)

	return nil
}

func LogoutUserService(c *gin.Context) error {
	// Get Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return errors.New("no token provided")
	}

	// Extract JWT token
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 {
		return errors.New("invalid token")
	}

	jwtToken := tokenParts[1]

	// Store token in Redis with expiration time (same as JWT expiry)
	ctx := context.Background()
	err := db.RedisClient.Set(ctx, jwtToken, "blacklisted", time.Hour*24).Err() // Expire in 24 hours
	if err != nil {
		return errors.New("failed to blacklist a token")
	}

	return nil
}