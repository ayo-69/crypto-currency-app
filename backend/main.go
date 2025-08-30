package main

import (
	"os"

	"github.com/ayo-69/crypto-price-tracker/backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	godotenv.Load()
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}

	routes.RegisterPriceRoutes(r)

	r.Run(":" + PORT)
}
