package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"goserv/controllers"
	"goserv/models"
	"net/http"
)

func JWTAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
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

		var user models.User
		err = models.GetUserByUsername(db, username, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
