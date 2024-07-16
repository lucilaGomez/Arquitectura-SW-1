package controllers

import (
	"log"
	"net/http"
	"proyecto/dtos"
	"proyecto/models"
	"proyecto/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateReservation handles the creation of a new reservation
func CreateReservation(c *gin.Context) {
	var dto dtos.ReservationDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parsear las fechas
	checkIn, err := time.Parse("2006-01-02", dto.CheckIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fecha check-in inválida"})
		return
	}

	checkOut, err := time.Parse("2006-01-02", dto.CheckOut)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fecha check-out inválida"})
		return
	}

	// Verificar disponibilidad
	if err := services.CheckAvailability(dto.HotelID, checkIn, checkOut); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear reserva
	reservation := models.Reservation{
		HotelID:  dto.HotelID,
		UserID:   dto.UserID,
		CheckIn:  checkIn,
		CheckOut: checkOut,
	}

	if err := services.CreateReservation(&reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}

	c.JSON(http.StatusCreated, reservation)
}

// GetReservations handles fetching all reservations
func GetReservations(c *gin.Context) {
	reservations, err := services.GetReservations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

// GetUserReservations handles fetching reservations of the logged-in user
func GetUserReservations(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: no user in context"})
		return
	}

	log.Printf("User ID: %d", user.(models.User).ID) // Log del ID del usuario

	reservations, err := services.GetUserReservations(user.(models.User).ID)
	if err != nil {
		log.Printf("Error fetching reservations: %v", err) // Log del error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}

// GetReservation handles fetching a single reservation by ID
func GetReservation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	reservation, err := services.GetReservationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

// UpdateReservation handles updating a reservation's details
func UpdateReservation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var dto dtos.ReservationDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation, err := services.UpdateReservation(id, &dto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found or update failed"})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

// DeleteReservation handles deleting a reservation
func DeleteReservation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteReservation(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reservation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted"})
}

func GetMyReservations(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	reservations, err := services.GetUserReservations(user.(models.User).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
