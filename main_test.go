package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootRoute(t *testing.T) {
	router := setupRouter()

	respRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(respRecorder, request)

	assert.Equal(t, http.StatusOK, respRecorder.Code)
	assert.Equal(t, "{\"message\":\"Go Prime API v0.0\"}", respRecorder.Body.String())
}

func TestNumberInputRouteWhenInputIsValidNumber(t *testing.T) {
	router := setupRouter()

	respRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/55", nil)
	router.ServeHTTP(respRecorder, request)

	assert.Equal(t, http.StatusOK, respRecorder.Code)
	assert.Equal(t, "53", respRecorder.Body.String())
}

func TestNumberInputRouteWhenInputIsInvalidNumber(t *testing.T) {
	router := setupRouter()

	respRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/invalid", nil)
	router.ServeHTTP(respRecorder, request)

	assert.Equal(t, http.StatusUnprocessableEntity, respRecorder.Code)
}