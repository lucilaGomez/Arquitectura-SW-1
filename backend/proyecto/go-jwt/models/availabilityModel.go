// models/availabilityModel.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Availability struct {
	gorm.Model
	HotelID   uint      `json:"hotel_id"`
	Date      time.Time `json:"date"`
	Available int       `json:"available"`
}
