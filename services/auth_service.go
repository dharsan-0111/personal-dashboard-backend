package services

import (
	"errors"
	"personal-dashboard-backend/api/user/model"
	"personal-dashboard-backend/db"

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
	var user model.User
	var storedPassword string

	if err := c.ShouldBindJSON(&user); err != nil {
		return err
	}

	err := db.DB.QueryRow("SELECT password FROM users WHERE email = $1", user.Email).Scan(&storedPassword)
	if err != nil { 
		return errors.New("user not found")
	}

	if !CheckPasswordHash(user.Password, storedPassword) {
		return errors.New("invalid password")
	}

	return nil
}