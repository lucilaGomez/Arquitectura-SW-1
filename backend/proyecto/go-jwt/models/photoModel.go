package models

type Photo struct {
	ID      uint   `gorm:"primaryKey"`
	URL     string `json:"url"`
	HotelID uint   `json:"hotel_id"`
}
