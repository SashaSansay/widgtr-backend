package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"widgtr-backend/internal/middlewares"
	"widgtr-backend/internal/models"
	"widgtr-backend/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDB()

	defer func() {
		models.DisconnectDB()
	}()

	r := gin.Default()
	middlewares.SetupMiddlewares(r)
	routes.SetupRoutes(r)
	r.Run()
}
