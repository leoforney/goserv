package config

import (
	"github.com/gin-contrib/cors"
)

const JWTSecretKey = "supersecretkeypleasechangeme"

func CORSConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}
}
