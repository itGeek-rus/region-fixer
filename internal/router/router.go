package router

import (
	"region-fixer/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.HealthHandler)

	return router
}
