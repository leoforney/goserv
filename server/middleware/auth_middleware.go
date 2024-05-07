package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"goserv/controllers"
	"goserv/models"
	"net/http"
)

func JWTAuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed JWT"})
			c.Abort()
			return
		}

		claims, err := controllers.ValidateJWT(authHeader[7:])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired JWT"})
			c.Abort()
			return
		}

		username := (*claims)["username"].(string)

		user, err := models.GetUserByUsername(db, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", *user)
		c.Next()
	}
}
