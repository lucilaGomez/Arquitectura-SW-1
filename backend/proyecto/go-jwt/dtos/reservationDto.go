package dtos

type ReservationDto struct {
	HotelID  uint   `json:"hotel_id" binding:"required"`
	UserID   uint   `json:"user_id" binding:"required"`
	CheckIn  string `json:"check_in" binding:"required"`
	CheckOut string `json:"check_out" binding:"required"`
}
