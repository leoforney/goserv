package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"goserv/config"
	"goserv/controllers"
	"goserv/middleware"
	"goserv/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	r := gin.Default()
	r.Use(cors.New(config.CORSConfig()))

	api := r.Group("/api")
	{
		api.POST("/register", func(c *gin.Context) {
			controllers.RegisterUser(c, db)
		})
		api.POST("/login", func(c *gin.Context) {
			controllers.Login(c, db)
		})

		authenticated := api.Group("/")
		authenticated.Use(middleware.JWTAuthMiddleware(db))
		{
			authenticated.GET("/me", controllers.MeEndpoint)
		}
	}

	if err = r.Run(); err != nil {
		panic(err)
	}
}
