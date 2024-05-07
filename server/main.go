package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"goserv/config"
	"goserv/controllers"
	"goserv/middleware"
	"goserv/models"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models.CreateTables(db)

	r := gin.Default()

	r.Use(cors.New(config.CORSConfig()))

	authenticated := r.Group("/")
	authenticated.Use(middleware.JWTAuthMiddleware(db))
	{
		authenticated.GET("/me", controllers.MeEndpoint)
	}

	r.POST("/register", func(c *gin.Context) {
		controllers.RegisterUser(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})

	if err = r.Run(); err != nil {
		return
	}
}
