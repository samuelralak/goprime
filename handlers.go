package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func rootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Go Prime API v0.0"})
}

func numberInputHandler(context *gin.Context) {
	validate := validator.New()
	number := context.Param("number")

	err := validate.Var(number, "required,numeric"); if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	message := "Your number is " + number
	context.JSON(http.StatusOK, gin.H{"message": message})
}