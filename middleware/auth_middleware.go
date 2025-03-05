package middleware

import (
	"context"
	"log"
	"net/http"
	"personal-dashboard-backend/db"
	"personal-dashboard-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) { 
		log.Println("Auth middleware")

		// get auth header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" { 
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// extract JWT token
		token := strings.Split(authHeader, " ")
		if len(token) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		jwtToken := token[1]

		// verify JWT token
		claims, err := utils.VerifyJWT(jwtToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		// check if token is blacklisted in Redis
		ctx := context.Background()
		_, redisErr := db.RedisClient.Get(ctx, jwtToken).Result()

		if redisErr == redis.Nil {
			// If key is NOT in Redis, it's a valid token (since we only store blacklisted tokens)
			log.Println("Token is valid, not found in blacklist")
		} else if redisErr != nil {
			// If there is an actual Redis error, return Internal Server Error
			log.Println("Error checking Redis:", redisErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		} else {
			// If token IS found in Redis, it is blacklisted
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been revoked"})
			c.Abort()
			return
		}

		// store user email in context
		c.Set("email", claims["email"])

		// continue if token is valid
		c.Next()
	}
}