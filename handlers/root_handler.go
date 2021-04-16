package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootHandler(context *gin.Context) {
	context.String(http.StatusOK, "Go Prime API v0.0")
}
