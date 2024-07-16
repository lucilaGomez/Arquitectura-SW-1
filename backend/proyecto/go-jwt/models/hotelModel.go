// models/hotelModel.go
package models

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	//Id          int       `gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Address     string    `json:"address" binding:"required"`
	City        string    `json:"city" binding:"required"`
	Country     string    `json:"country" binding:"required"`
	Amenities   []Amenity `gorm:"many2many:hotel_amenities" json:"amenities"`
	Photos      []Photo   `json:"photos"`
}
