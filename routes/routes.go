package routes

import "github.com/gin-gonic/gin"

func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	
	AuthRoutes(r)
	UserRoutes(r)

	return r
}