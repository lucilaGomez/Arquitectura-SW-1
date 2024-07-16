package services

import (
	"errors"
	"proyecto/dtos"
	"proyecto/initializers"
	"proyecto/models"
	"time"
)

// CreateReservation creates a new reservation in the database
func CreateReservation(reservation *models.Reservation) error {
	checkIn := reservation.CheckIn
	checkOut := reservation.CheckOut

	var availabilities []models.Availability
	if err := initializers.DB.Where("hotel_id = ? AND date >= ? AND date < ?", reservation.HotelID, checkIn, checkOut).Find(&availabilities).Error; err != nil {
		return err
	}

	for _, availability := range availabilities {
		if availability.Available <= 0 {
			return errors.New("no availability")
		}
	}

	if err := initializers.DB.Create(reservation).Error; err != nil {
		return err
	}

	// Actualizar disponibilidad
	for _, availability := range availabilities {
		availability.Available--
		if err := initializers.DB.Save(&availability).Error; err != nil {
			return err
		}
	}

	return nil
}

// GetReservations retrieves all reservations from the database
func GetReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	if err := initializers.DB.Preload("Hotel").Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetUserReservations(userID uint) ([]models.Reservation, error) {
	var reservations []models.Reservation
	if err := initializers.DB.Where("user_id = ?", userID).Preload("Hotel").Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

// GetReservationByID retrieves a single reservation by its ID from the database
func GetReservationByID(id int) (models.Reservation, error) {
	var reservation models.Reservation
	if err := initializers.DB.Preload("Hotel").First(&reservation, id).Error; err != nil {
		return models.Reservation{}, err
	}
	return reservation, nil
}

// UpdateReservation updates an existing reservation in the database
func UpdateReservation(id int, dto *dtos.ReservationDto) (models.Reservation, error) {
	var reservation models.Reservation
	if err := initializers.DB.First(&reservation, id).Error; err != nil {
		return models.Reservation{}, err
	}

	// Parse dates
	checkIn, err := time.Parse("2006-01-02", dto.CheckIn)
	if err != nil {
		return models.Reservation{}, err
	}
	checkOut, err := time.Parse("2006-01-02", dto.CheckOut)
	if err != nil {
		return models.Reservation{}, err
	}

	// Update fields
	reservation.HotelID = dto.HotelID
	reservation.UserID = dto.UserID
	reservation.CheckIn = checkIn
	reservation.CheckOut = checkOut

	if err := initializers.DB.Save(&reservation).Error; err != nil {
		return models.Reservation{}, err
	}

	return reservation, nil
}

// DeleteReservation deletes a reservation from the database
func DeleteReservation(id int) error {
	var reservation models.Reservation
	if err := initializers.DB.First(&reservation, id).Error; err != nil {
		return err
	}

	if err := initializers.DB.Delete(&reservation).Error; err != nil {
		return err
	}

	// Update availability
	if err := updateAvailability(reservation.HotelID, reservation.CheckIn, reservation.CheckOut, 1); err != nil {
		return err
	}

	return nil
}

// CheckAvailability checks the availability of a hotel between two dates
func CheckAvailability(hotelID uint, checkIn, checkOut time.Time) error {
	var availabilities []models.Availability
	if err := initializers.DB.Where("hotel_id = ? AND date >= ? AND date < ?", hotelID, checkIn, checkOut).Find(&availabilities).Error; err != nil {
		return err
	}

	for _, availability := range availabilities {
		if availability.Available <= 0 {
			return errors.New("no availability")
		}
	}
	return nil
}

// updateAvailability updates the availability of a hotel between two dates
func updateAvailability(hotelID uint, checkIn, checkOut time.Time, quantity int) error {
	var availabilities []models.Availability
	if err := initializers.DB.Where("hotel_id = ? AND date >= ? AND date < ?", hotelID, checkIn, checkOut).Find(&availabilities).Error; err != nil {
		return err
	}

	for _, availability := range availabilities {
		availability.Available += quantity
		if err := initializers.DB.Save(&availability).Error; err != nil {
			return err
		}
	}

	return nil
}
