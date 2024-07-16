package controllers

import (
	"bytes"
	"encoding/json"

	//	"errors"
	"net/http"
	"net/http/httptest"
	"proyecto/dtos"
	"proyecto/models"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

/*
This is a unit test function for the GetHotelsWithAvailability endpoint in a Gin web application.
It checks the behavior of the endpoint
when the request is missing required date parameters.
*
*/
func TestGetHotelsWithAvailability_MissingDates(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/hotels", GetHotelsWithAvailability)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hotels?start_date=2023-01-01", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Missing start_date or end_date query parameters
/*This function tests the GetAvailableHotels endpoint
to ensure it returns a 400 Bad Request
status when the end_date query parameter is missing.
*/
func TestMissingDateQueryParameters(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/hotels", GetAvailableHotels)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hotels?start_date=2023-01-01", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "start_date and end_date are required", response["error"])
}

// Mocking the service
type mockHotelService struct{}

func (m *mockHotelService) CreateHotel(dto dtos.HotelDto) (*models.Hotel, error) {
	return &models.Hotel{
		// ID:      1,
		Name:        dto.Name,
		Description: dto.Description,
		Address:     dto.Address,
		City:        dto.City,
		Country:     dto.Country,
	}, nil
}

/*
This function tests the successful creation of a
hotel using a mock service and the Gin framework.
*/
func TestCreateHotelSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &mockHotelService{}
	controller := NewHotelController(mockService)

	router := gin.Default()
	router.POST("/hotels", controller.CreateHotel)

	hotelDto := dtos.HotelDto{
		Name:        "Test Hotel",
		Description: "Test hotel description",
		Address:     "123 Test St",
		City:        "Test City",
		Country:     "Test Country",
	}

	jsonValue, _ := json.Marshal(hotelDto)
	req, _ := http.NewRequest(http.MethodPost, "/hotels", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

/*
This function is a unit test for the CreateHotel endpoint, specifically testing
the scenario where an invalid JSON payload is sent.
*/
// Replicate the original test function 'TestCreateHotelInvalidJSON'
func TestReplicateCreateHotelInvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/hotels", CreateHotel)

	invalidJson := `{"Name": "Test Hotel", "Address": 123}`

	req, _ := http.NewRequest(http.MethodPost, "/hotels", bytes.NewBufferString(invalidJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Handling invalid hotel ID in the request parameter
/*
This code is a unit test function for testing the UpdateHotel handler in a Gin web application. It specifically tests the scenario where an
invalid hotel ID is provided in the request URL.
*/
func TestUpdateHotel_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/hotels/:id", UpdateHotel)

	hotelDto := dtos.HotelDto{
		Name:    "Updated Hotel",
		Address: "123 Updated St",
		//Rating: 5,
	}

	jsonValue, _ := json.Marshal(hotelDto)
	req, _ := http.NewRequest(http.MethodPut, "/hotels/invalid", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// ___________________________________________________
// Successfully updating a hotel with valid ID and data
func TestUpdateHotel_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/hotels/:id", UpdateHotel)

	hotelDto := dtos.HotelDto{
		Name:    "Updated Hotel",
		Address: "123 Updated St",
		//Rating: 5,
	}

	jsonValue, _ := json.Marshal(hotelDto)
	req, _ := http.NewRequest(http.MethodPut, "/hotels/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

//___________________________________________________
