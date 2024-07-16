// initializers/syncDatabase.go
package initializers

import "proyecto/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Hotel{})
	DB.AutoMigrate(&models.Reservation{})
	DB.AutoMigrate(&models.Photo{})
	DB.AutoMigrate(&models.Amenity{})
	DB.AutoMigrate(&models.Availability{}) // Añade esta línea
}
