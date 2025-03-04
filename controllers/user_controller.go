package controllers

import (
	"net/http"
	"personal-dashboard-backend/services"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	identifier := c.Query("username")
	if identifier == "" { 
		identifier = c.Query("email")
	}

	if identifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Identifier (email or username) is required"})
		return
	}

	userData, err := services.GetUserService(identifier)
	if err != nil { 
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userData})
}