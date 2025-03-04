package controllers

import (
	"log"
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
	log.Println("Login")
}
