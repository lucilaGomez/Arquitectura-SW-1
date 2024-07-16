package models

import (
	"gorm.io/gorm"
)

type Amenity struct {
	gorm.Model
	Name string `json:"name"`
}
