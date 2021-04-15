package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Go Prime API v0.0"})
	})

	router.GET("/:number", func(context *gin.Context) {
		number := context.Param("number")
		message := "Your number is " + number
		context.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.Run()
}
