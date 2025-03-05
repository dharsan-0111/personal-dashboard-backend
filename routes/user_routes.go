package routes

import (
	"personal-dashboard-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) { 
	r.GET("", controllers.GetUser)
}