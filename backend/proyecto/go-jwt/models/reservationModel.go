package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	HotelID  uint      `json:"hotel_id"`
	UserID   uint      `json:"user_id"`
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
	Hotel    Hotel     `gorm:"foreignKey:HotelID"`
}
