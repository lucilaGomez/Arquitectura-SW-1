package main

import (
	"proyecto/controllers"
	"proyecto/initializers"
	"proyecto/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// Configuración CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"}, // Cambia esto por el origen correcto de tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rutas y controladores
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/auth/me", middleware.RequireAuth, controllers.GetCurrentUser)
	r.POST("/auth/logout", controllers.Logout)

	// Hotels
	r.POST("/hotels", middleware.RequireAuth, middleware.RequireAdmin, controllers.CreateHotel)
	r.GET("/hotels", controllers.GetHotels)
	r.GET("/hotels/:id", controllers.GetHotel)
	r.GET("/available-hotels", controllers.GetAvailableHotels)
	r.GET("/hotels-availability", controllers.GetHotelsWithAvailability)
	r.PUT("/hotels/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.UpdateHotel)
	r.DELETE("/hotels/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.DeleteHotel)

	// Photos
	r.POST("/photos", middleware.RequireAuth, middleware.RequireAdmin, controllers.CreatePhoto)
	r.GET("/photos", controllers.GetPhotos)
	r.GET("/photos/:id", controllers.GetPhoto)
	r.PUT("/photos/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.UpdatePhoto)
	r.DELETE("/photos/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.DeletePhoto)

	// Amenities
	r.POST("/amenities", middleware.RequireAuth, middleware.RequireAdmin, controllers.CreateAmenity)
	r.GET("/amenities", controllers.GetAllAmenities)
	r.GET("/amenities/:id", controllers.GetAmenityByID)
	r.PUT("/amenities/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.UpdateAmenity)
	r.DELETE("/amenities/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.DeleteAmenity)

	// Reservations
	r.POST("/reservations", middleware.RequireAuth, controllers.CreateReservation)
	r.GET("/reservations", middleware.RequireAuth, controllers.GetReservations)
	r.GET("/reservations/:id", middleware.RequireAuth, controllers.GetReservation)
	r.PUT("/reservations/:id", middleware.RequireAuth, controllers.UpdateReservation)
	r.DELETE("/reservations/:id", middleware.RequireAuth, controllers.DeleteReservation)
	r.GET("/reservations/user", middleware.RequireAuth, controllers.GetUserReservations)
	r.GET("/reservations/my", middleware.RequireAuth, controllers.GetMyReservations)

	// Availability
	r.GET("/availability", middleware.RequireAuth, middleware.RequireAdmin, controllers.GetAvailability)
	r.POST("/availability", middleware.RequireAuth, middleware.RequireAdmin, controllers.CreateInitialAvailability)
	r.PUT("/availability", middleware.RequireAuth, middleware.RequireAdmin, controllers.UpdateAvailability)
	r.DELETE("/availability/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.DeleteAvailability)

	r.Run()
}
