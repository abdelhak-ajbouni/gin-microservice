package main

import (
	"log"
	"os"

	"github.com/abdelhak-ajbouni/gin-microservice/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set gin to release mode
	mode := os.Getenv("MODE")
	if mode == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/health", routers.GetHealth)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Endpoint Not Found",
		})
	})

	// Listen and Serve
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
