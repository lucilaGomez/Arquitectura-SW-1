package controllers

import (
	"bytes"
	//"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	//"time"

	//"proyecto/dtos"
	//"proyecto/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
)

// Returns a bad request status when input JSON is invalid
func TestCreateReservation_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/reservations", CreateReservation)

	invalidJSON := `{"HotelID": "invalid", "UserID": 1, "CheckIn": "2023-10-01", "CheckOut": "2023-10-10"}`

	req, _ := http.NewRequest(http.MethodPost, "/reservations", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

//__________________________________________
// MockService es un mock del servicio que usa CreateReservation.
