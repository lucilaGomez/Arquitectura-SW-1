package dtos

type HotelDto struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Address     string   `json:"address" binding:"required"`
	City        string   `json:"city" binding:"required"`
	Country     string   `json:"country" binding:"required"`
	Amenities   []string `json:"amenities"`
	Photos      []string `json:"photos"`
}
