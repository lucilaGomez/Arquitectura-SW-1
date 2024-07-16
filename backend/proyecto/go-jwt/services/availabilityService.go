package services

import (
	"proyecto/dtos"
	"proyecto/initializers"
	"proyecto/models"
	"time"
)

func GetAvailability(hotelID string, startDate, endDate time.Time) ([]models.Availability, error) {
	var availabilities []models.Availability
	query := initializers.DB.Where("hotel_id = ? AND date >= ? AND date <= ?", hotelID, startDate, endDate)
	if err := query.Find(&availabilities).Error; err != nil {
		return nil, err
	}
	return availabilities, nil
}

func CreateInitialAvailability(dto *dtos.InitialAvailabilityDto) error {
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)

	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		availability := models.Availability{
			HotelID:   dto.HotelID,
			Date:      d,
			Available: dto.Available,
		}

		if err := initializers.DB.Create(&availability).Error; err != nil {
			return err
		}
	}

	return nil
}

func UpdateAvailability(dto *dtos.UpdateAvailabilityDto) error {
	checkInParsed, err := time.Parse("2006-01-02", dto.CheckIn)
	if err != nil {
		return err
	}
	checkOutParsed, err := time.Parse("2006-01-02", dto.CheckOut)
	if err != nil {
		return err
	}

	var availabilities []models.Availability
	query := initializers.DB.Where("hotel_id = ? AND date >= ? AND date < ?", dto.HotelID, checkInParsed, checkOutParsed)
	if err := query.Find(&availabilities).Error; err != nil {
		return err
	}

	for _, availability := range availabilities {
		availability.Available = dto.Quantity
		if err := initializers.DB.Save(&availability).Error; err != nil {
			return err
		}
	}

	return nil
}

func DeleteAvailability(id string) error {
	if err := initializers.DB.Delete(&models.Availability{}, id).Error; err != nil {
		return err
	}
	return nil
}
