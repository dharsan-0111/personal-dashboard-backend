package controllers

import (
	"net/http"
	"personal-dashboard-backend/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	err := services.RegisterUserService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	err := services.LoginUserService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}

func Logout(c *gin.Context) {
	err := services.LogoutUserService(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occured while logging out"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
