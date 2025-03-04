package routes

import (
	"personal-dashboard-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) { 
	authGroup := r.Group("/auth")

	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/login", controllers.Login)
}