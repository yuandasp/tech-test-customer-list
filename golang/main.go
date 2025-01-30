package main

import (
	"log"
	"os"

	"tech-test/config"
	"tech-test/models"
	"tech-test/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(cors.Default())
	
	config.ConnectDB()
	
	config.DB.AutoMigrate(&models.Customer{})

	routes.CustomerRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
   		port = "8080"
	}
	r.Run(":" + port)
}
