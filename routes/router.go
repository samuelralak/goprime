package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelralak/goprime/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Root route
	// Returns message with API version
	router.GET("/", handlers.RootHandler)

	// Number route takes in a number as a parameter
	// Returns the highest number lower than the inout number
	router.GET("/:number", handlers.NumberInputHandler)

	return router
}