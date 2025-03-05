package routes

import (
	"personal-dashboard-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	
	AuthRoutes(r)

	protected := r.Group("/user")
	protected.Use(middleware.AuthMiddleware())
	{
		UserRoutes(protected)
	}

	return r
}