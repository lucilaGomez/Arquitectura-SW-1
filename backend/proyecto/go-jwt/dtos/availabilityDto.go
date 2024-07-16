// dtos/availabilityDto.go
package dtos

type InitialAvailabilityDto struct {
	HotelID   uint `json:"hotel_id" binding:"required"`
	Available int  `json:"available" binding:"required"`
}

type UpdateAvailabilityDto struct {
	HotelID  uint   `json:"hotel_id" binding:"required"`
	CheckIn  string `json:"check_in" binding:"required"`  // Fecha en formato "YYYY-MM-DD"
	CheckOut string `json:"check_out" binding:"required"` // Fecha en formato "YYYY-MM-DD"
	Quantity int    `json:"quantity" binding:"required"`  // Cantidad de habitaciones a reservar
}
