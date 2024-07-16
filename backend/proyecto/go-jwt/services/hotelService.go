package services

import (
	"fmt"
	"proyecto/dtos"
	"proyecto/initializers"
	"proyecto/models"
	"time"
)

func GetHotelsWithAvailability(startDate time.Time, endDate time.Time) ([]models.Hotel, error) {
	var hotels []models.Hotel
	var availableHotels []models.Hotel

	initializers.DB.Preload("Amenities").Preload("Photos").Find(&hotels)

	for _, hotel := range hotels {
		var totalAvailability int64
		initializers.DB.Model(&models.Availability{}).
			Where("hotel_id = ? AND date >= ? AND date <= ? AND available > 0", hotel.ID, startDate, endDate).
			Count(&totalAvailability)

		if totalAvailability == int64(endDate.Sub(startDate).Hours()/24)+1 {
			availableHotels = append(availableHotels, hotel)
		}
	}

	return availableHotels, nil
}

func GetAvailableHotels(startDate time.Time, endDate time.Time) ([]models.Hotel, error) {
	var hotels []models.Hotel
	var availableHotels []models.Hotel

	initializers.DB.Preload("Amenities").Preload("Photos").Find(&hotels)

	for _, hotel := range hotels {
		var availabilities []models.Availability
		initializers.DB.Where("hotel_id = ? AND date >= ? AND date <= ? AND available > 0", hotel.ID, startDate, endDate).Find(&availabilities)

		if len(availabilities) == int(endDate.Sub(startDate).Hours()/24)+1 {
			availableHotels = append(availableHotels, hotel)
		}
	}

	return availableHotels, nil
}

func CreateHotel(hotelDto dtos.HotelDto) (*models.Hotel, error) {
	// Check if a hotel with the same Address already exists
	var count int64
	if err := initializers.DB.Model(&models.Hotel{}).
		Where("address = ?", hotelDto.Address).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, fmt.Errorf("another hotel with the same Address already exists")
	}

	// Check if the combination of Name, Address, City, and Country already exists
	var anotherCount int64
	if err := initializers.DB.Model(&models.Hotel{}).
		Where("name != ?", hotelDto.Name).
		Where("address = ?", hotelDto.Address).
		Where("city = ?", hotelDto.City).
		Where("country = ?", hotelDto.Country).
		Count(&anotherCount).Error; err != nil {
		return nil, err
	}

	if anotherCount > 0 {
		return nil, fmt.Errorf("hotel with the same Address and different Name already exists")
	}

	var amenities []models.Amenity
	var photos []models.Photo

	for _, name := range hotelDto.Amenities {
		var amenity models.Amenity
		if err := initializers.DB.Where("name = ?", name).First(&amenity).Error; err != nil {
			return nil, err
		}
		amenities = append(amenities, amenity)
	}

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

	if err := initializers.DB.Create(&hotel).Error; err != nil {
		return nil, err
	}

	return &hotel, nil
}

func GetHotels() ([]models.Hotel, error) {
	var hotels []models.Hotel
	if err := initializers.DB.Preload("Amenities").Preload("Photos").Find(&hotels).Error; err != nil {
		return nil, err
	}
	return hotels, nil
}

func GetHotel(id int) (*models.Hotel, error) {
	var hotel models.Hotel
	if err := initializers.DB.Preload("Amenities").Preload("Photos").First(&hotel, id).Error; err != nil {
		return nil, err
	}
	return &hotel, nil
}

func UpdateHotel(id int, hotelDto dtos.HotelDto) (*models.Hotel, error) {
	var hotel models.Hotel
	if err := initializers.DB.Preload("Amenities").Preload("Photos").First(&hotel, id).Error; err != nil {
		return nil, err
	}

	var amenities []models.Amenity
	for _, name := range hotelDto.Amenities {
		var amenity models.Amenity
		if err := initializers.DB.Where("name = ?", name).First(&amenity).Error; err != nil {
			return nil, err
		}
		amenities = append(amenities, amenity)
	}

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

	if err := initializers.DB.Save(&hotel).Error; err != nil {
		return nil, err
	}

	return &hotel, nil
}

func DeleteHotel(id int) error {
	if err := initializers.DB.Delete(&models.Hotel{}, id).Error; err != nil {
		return err
	}
	return nil
}

// MOCK
type HotelService interface {
	CreateHotel(dto dtos.HotelDto) (*models.Hotel, error)
}

// Implementación real del servicio, si ya existe
type hotelService struct{}

func NewHotelService() HotelService {
	return &hotelService{}
}

func (s *hotelService) CreateHotel(dto dtos.HotelDto) (*models.Hotel, error) {
	// Aquí va la lógica para crear un hotel
	return nil, nil
}

/*
//___________________________________________________________________
type HotelServiceU interface {
    CreateHotel(dto dtos.HotelDto) (*models.Hotel, error)
    UpdateHotel(id int, dto dtos.HotelDto) (*models.Hotel, error)
}

type hotelServiceU struct{}

func NewHotelServiceU() HotelService {
    return &hotelService{}
}

func (s *hotelService) CreateHotelU(dto dtos.HotelDto) (*models.Hotel, error) {
    // Implementación real
    return nil, nil
}

func (s *hotelService) UpdateHotel(id int, dto dtos.HotelDto) (*models.Hotel, error) {
    // Implementación real
    return nil, nil
}
*/
