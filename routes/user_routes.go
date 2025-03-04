package routes

import (
	"personal-dashboard-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) { 
	userGroup := r.Group("/user")

	userGroup.GET("", controllers.GetUser)
}