package controllers

import (
	"net/http"
	"proyecto/dtos"
	"proyecto/services"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetHotelsWithAvailability(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}

	availableHotels, err := services.GetHotelsWithAvailability(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch available hotels"})
		return
	}

	c.JSON(http.StatusOK, availableHotels)
}

func GetAvailableHotels(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}

	availableHotels, err := services.GetAvailableHotels(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch available hotels"})
		return
	}

	c.JSON(http.StatusOK, availableHotels)
}

func CreateHotel(c *gin.Context) {
	var hotelDto dtos.HotelDto
	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel, err := services.CreateHotel(hotelDto)
	if err != nil {
		// Check if the error is due to duplicate combination
		if strings.Contains(err.Error(), "hotel with the same Name, Address, City, and Country already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hotel"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ID": hotel.ID})
}

func GetHotels(c *gin.Context) {
	hotels, err := services.GetHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hotels"})
		return
	}

	c.JSON(http.StatusOK, hotels)
}

func GetHotel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	hotel, err := services.GetHotel(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

func UpdateHotel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dtos.HotelDto
	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel, err := services.UpdateHotel(id, hotelDto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found or update failed"})
		return
	}

	c.JSON(http.StatusOK, hotel) // Devuelve el hotel actualizado correctamente
}

func DeleteHotel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	if err := services.DeleteHotel(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted"})
}

/*
func GetHotelsWithAvailability(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}

	var hotels []models.Hotel
	initializers.DB.Preload("Amenities").Preload("Photos").Find(&hotels)

	var availableHotels []models.Hotel

	for _, hotel := range hotels {
		var totalAvailability int64
		initializers.DB.Model(&models.Availability{}).
			Where("hotel_id = ? AND date >= ? AND date <= ? AND available > 0", hotel.ID, startDate, endDate).
			Count(&totalAvailability)

		if totalAvailability == int64(endDate.Sub(startDate).Hours()/24)+1 {
			availableHotels = append(availableHotels, hotel)
		}
	}

	c.JSON(http.StatusOK, availableHotels)
}

func GetAvailableHotels(c *gin.Context) {
	var hotels []models.Hotel
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Println("Error parsing start_date:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date"})
		return
	}

	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Println("Error parsing end_date:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date"})
		return
	}

	var availableHotels []models.Hotel

	initializers.DB.Preload("Amenities").Preload("Photos").Find(&hotels)

	for _, hotel := range hotels {
		var availabilities []models.Availability
		initializers.DB.Where("hotel_id = ? AND date >= ? AND date <= ? AND available > 0", hotel.ID, startDateParsed, endDateParsed).Find(&availabilities)

		if len(availabilities) == int(endDateParsed.Sub(startDateParsed).Hours()/24)+1 {
			availableHotels = append(availableHotels, hotel)
		}
	}

	c.JSON(http.StatusOK, availableHotels)
}

func CreateHotel(c *gin.Context) {
	var hotelDto dtos.HotelDto
	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar y agregar solo amenities existentes
	var amenities []models.Amenity
	for _, name := range hotelDto.Amenities {
		var amenity models.Amenity
		if err := initializers.DB.Where("name = ?", name).First(&amenity).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Amenity " + name + " does not exist"})
			return
		}
		amenities = append(amenities, amenity)
	}

	// Convert DTO to Model
	var photos []models.Photo
	for _, url := range hotelDto.Photos {
		photos = append(photos, models.Photo{URL: url})
	}

	hotel := models.Hotel{
		Name:        hotelDto.Name,
		Description: hotelDto.Description,
		Address:     hotelDto.Address,
		City:        hotelDto.City,
		Country:     hotelDto.Country,
		Amenities:   amenities,
		Photos:      photos,
	}

	initializers.DB.Create(&hotel)
	c.JSON(http.StatusCreated, hotel)
}

func GetHotels(c *gin.Context) {
	var hotels []models.Hotel
	initializers.DB.Preload("Amenities").Preload("Photos").Find(&hotels)
	c.JSON(http.StatusOK, hotels)
}

func GetHotel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hotel models.Hotel
	if err := initializers.DB.Preload("Amenities").Preload("Photos").First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}
	c.JSON(http.StatusOK, hotel)
}

func UpdateHotel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hotel models.Hotel
	if err := initializers.DB.Preload("Amenities").Preload("Photos").First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	var hotelDto dtos.HotelDto
	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar y agregar solo amenities existentes
	var amenities []models.Amenity
	for _, name := range hotelDto.Amenities {
		var amenity models.Amenity
		if err := initializers.DB.Where("name = ?", name).First(&amenity).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Amenity " + name + " does not exist"})
			return
		}
		amenities = append(amenities, amenity)
	}

	// Convert DTO to Model
	var photos []models.Photo
	for _, url := range hotelDto.Photos {
		photos = append(photos, models.Photo{URL: url})
	}

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.Address = hotelDto.Address
	hotel.City = hotelDto.City
	hotel.Country = hotelDto.Country
	hotel.Amenities = amenities
	hotel.Photos = photos

	initializers.DB.Save(&hotel)
	c.JSON(http.StatusOK, hotel)
}

func DeleteHotel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := initializers.DB.Delete(&models.Hotel{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted"})
}
*/

//MOCK

type HotelController struct {
	hotelService services.HotelService
}

func NewHotelController(service services.HotelService) *HotelController {
	return &HotelController{
		hotelService: service,
	}
}

func (ctrl *HotelController) CreateHotel(c *gin.Context) {
	var hotelDto dtos.HotelDto
	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel, err := ctrl.hotelService.CreateHotel(hotelDto)
	if err != nil {
		// Check if the error is due to duplicate combination
		if strings.Contains(err.Error(), "hotel with the same Name, Address, City, and Country already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hotel"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ID": hotel.ID})
}

/*
//___________________________________________________________________________________
type HotelControllerU struct {
    hotelService services.HotelService
}

func NewHotelControllerU(service services.HotelService) *HotelController {
    return &HotelController{
        hotelService: service,
    }
}

func (ctrl *HotelController) CreateHotelU(c *gin.Context) {
    var hotelDto dtos.HotelDto
    if err := c.ShouldBindJSON(&hotelDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hotel, err := ctrl.hotelService.CreateHotel(hotelDto)
    if err != nil {
        if strings.Contains(err.Error(), "hotel with the same Name, Address, City, and Country already exists") {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hotel"})
        }
        return
    }

    c.JSON(http.StatusCreated, gin.H{"ID": hotel.ID})
}

func (ctrl *HotelController) UpdateHotel(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var hotelDto dtos.HotelDto
    if err := c.ShouldBindJSON(&hotelDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hotel, err := ctrl.hotelService.UpdateHotel(id, hotelDto)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found or update failed"})
        return
    }

    c.JSON(http.StatusOK, hotel)
}
*/
