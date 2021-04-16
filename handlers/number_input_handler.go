package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"math"
	"net/http"
	"strconv"
)

// handles number input request and
// returns highest prime number lower than the input
func NumberInputHandler(context *gin.Context) {
	validate := validator.New()
	number := context.Param("number")

	err := validate.Var(number, "required,numeric"); if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	numToInt, _ := strconv.Atoi(number)
	prime, err := getHighestPrime(numToInt); if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	context.String(http.StatusOK, "%d", prime)
}

// Uses sieve of Eratosthenes to
// Find the highest prime number lower than the limit
func getHighestPrime(limit int) (int, error) {
	var result int

	primeStore := make([]bool, limit)
	limitSqrt := int(math.Sqrt(float64(limit)))

	for i := 2; i <= limitSqrt; i++ {
		for j := i * i; j < limit; j += i {
			primeStore[j] = true
		}
	}

	for i := len(primeStore) - 1; i >= 2; i-- {
		if primeStore[i] == false {
			return i, nil
		}
	}

	return result, errors.New("no prime number found")
}