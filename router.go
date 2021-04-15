package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Root route
	// Returns message with API version
	router.GET("/", rootHandler)

	// Number route takes in a number as a parameter
	// Returns the highest number lower than the inout number
	router.GET("/:number", numberInputHandler)

	return router
}