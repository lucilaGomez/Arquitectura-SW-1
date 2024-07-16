package controllers

import (
	"net/http"
	"proyecto/dtos"
	"proyecto/services"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAvailability(c *gin.Context) {
	hotelID := c.Query("hotel_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if hotelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}

	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}

	availabilities, err := services.GetAvailability(hotelID, startDateParsed, endDateParsed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, availabilities)
}

func CreateInitialAvailability(c *gin.Context) {
	var dto dtos.InitialAvailabilityDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateInitialAvailability(&dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Initial availability created successfully"})
}

func UpdateAvailability(c *gin.Context) {
	var dto dtos.UpdateAvailabilityDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateAvailability(&dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Availability updated successfully"})
}

func DeleteAvailability(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteAvailability(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Availability deleted"})
}

/*func GetAvailability(c *gin.Context) {
	hotelID := c.Query("hotel_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if hotelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}

	// Parsear las fechas
	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}

	var availabilities []models.Availability
	query := initializers.DB.Where("hotel_id = ? AND date >= ? AND date <= ?", hotelID, startDateParsed, endDateParsed)
	query.Find(&availabilities)

	c.JSON(http.StatusOK, availabilities)
}

func CreateInitialAvailability(c *gin.Context) {
	var dto dtos.InitialAvailabilityDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate := time.Date(2004, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2030, 12, 31, 0, 0, 0, 0, time.UTC)

	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		availability := models.Availability{
			HotelID:   dto.HotelID,
			Date:      d,
			Available: dto.Available,
		}

		if err := initializers.DB.Create(&availability).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create availability"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Initial availability created successfully"})
}

func UpdateAvailability(c *gin.Context) {
	var dto dtos.UpdateAvailabilityDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parsear las fechas
	checkInParsed, err := time.Parse("2006-01-02", dto.CheckIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid check_in format"})
		return
	}
	checkOutParsed, err := time.Parse("2006-01-02", dto.CheckOut)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid check_out format"})
		return
	}

	var availabilities []models.Availability
	initializers.DB.Where("hotel_id = ? AND date >= ? AND date < ?", dto.HotelID, checkInParsed, checkOutParsed).Find(&availabilities)

	for _, availability := range availabilities {
		availability.Available -= dto.Quantity
		initializers.DB.Save(&availability)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Availability updated successfully"})
}

func DeleteAvailability(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.Availability{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "availability deleted"})
}*/
