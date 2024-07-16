package dtos

type PhotoDTO struct {
	ID      uint   `json:"id"`
	URL     string `json:"url" binding:"required"`
	HotelID uint   `json:"hotel_id" binding:"required"`
}

type PhotosDTO []PhotoDTO
