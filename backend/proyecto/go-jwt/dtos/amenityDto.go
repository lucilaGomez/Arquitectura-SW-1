package dtos

type AmenityDTO struct {
	ID      uint   `json:"id"`
	Name    string `json:"name" binding:"required"`
	HotelID uint   `json:"hotel_id" binding:"required"`
}

type AmenitiesDTO []AmenityDTO
