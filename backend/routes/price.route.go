package routes

import (
	"github.com/ayo-69/crypto-price-tracker/backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPriceRoutes(router *gin.Engine) {
	router.GET("/price", controllers.GetPrice)
}
