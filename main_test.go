package main_test

import (
	"github.com/samuelralak/goprime/routes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var serverURL string

type responseBody struct {
	statusCode int
	body       []byte
}

func TestMain(m *testing.M) {
	router := routes.SetupRouter()
	server := httptest.NewServer(router)

	serverURL = server.URL
	defer server.Close()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRootRoute(t *testing.T) {
	respBody, err := makeRequest("/")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, http.StatusOK, respBody.statusCode)
	assert.Equal(t, "Go Prime API v0.0", string(respBody.body))
}

func TestNumberInputRouteWhenInputIsValidNumber(t *testing.T) {
	respBody, err := makeRequest("/55")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, http.StatusOK, respBody.statusCode)
	assert.Equal(t, "53", string(respBody.body))
}

func TestNumberInputRouteWhenInputIsNotNumber(t *testing.T) {
	respBody, err := makeRequest("/invalid")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, respBody.statusCode)
}

func TestNumberInputRouteWhenInputIsEqualOrLessThanTwo(t *testing.T) {
	respBody, err := makeRequest("/2")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, respBody.statusCode)
	assert.Equal(t, "no prime number found", string(respBody.body))
}

func makeRequest(url string) (responseBody, error) {
	resp, err := http.Get(serverURL + url)
	if err != nil {
		return responseBody{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseBody{}, err
	}

	result := responseBody{
		statusCode: resp.StatusCode,
		body:       body,
	}

	return result, nil
}
