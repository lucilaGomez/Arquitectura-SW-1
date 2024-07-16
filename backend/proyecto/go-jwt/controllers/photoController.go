package controllers

import (
	"net/http"
	"proyecto/models"
	"proyecto/services"

	"github.com/gin-gonic/gin"
)

// CreatePhoto maneja la creación de una nueva foto.
func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreatePhoto(&photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo created successfully"})
}

// GetPhotos maneja la obtención de todas las fotos.
func GetPhotos(c *gin.Context) {
	photos, err := services.GetAllPhotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

// GetPhoto maneja la obtención de una sola foto por su ID.
func GetPhoto(c *gin.Context) {
	id := c.Param("id")

	photo, err := services.GetPhotoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photo": photo})
}

// UpdatePhoto maneja la actualización de los detalles de una foto.
func UpdatePhoto(c *gin.Context) {
	id := c.Param("id")

	var updatedPhoto models.Photo
	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdatePhoto(id, &updatedPhoto); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

// DeletePhoto maneja la eliminación de una foto.
func DeletePhoto(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeletePhoto(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

/*
// CreatePhoto handles the creation of a new photo
func CreatePhoto(c *gin.Context) {
	var dto dtos.PhotoDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photo := models.Photo{URL: dto.URL, HotelID: dto.HotelID}
	if err := initializers.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo created successfully"})
}

// GetPhotos handles fetching all photos
func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	initializers.DB.Find(&photos)

	var dto dtos.PhotosDTO
	for _, photo := range photos {
		dto = append(dto, dtos.PhotoDTO{
			ID:      photo.ID,
			URL:     photo.URL,
			HotelID: photo.HotelID,
		})
	}

	c.JSON(http.StatusOK, gin.H{"photos": dto})
}

// GetPhoto handles fetching a single photo by ID
func GetPhoto(c *gin.Context) {
	id := c.Param("id")

	var photo models.Photo
	if err := initializers.DB.First(&photo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	dto := dtos.PhotoDTO{
		ID:      photo.ID,
		URL:     photo.URL,
		HotelID: photo.HotelID,
	}

	c.JSON(http.StatusOK, gin.H{"photo": dto})
}

// UpdatePhoto handles updating a photo's details
func UpdatePhoto(c *gin.Context) {
	id := c.Param("id")

	var dto dtos.PhotoDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var photo models.Photo
	if err := initializers.DB.First(&photo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	photo.URL = dto.URL
	initializers.DB.Save(&photo)

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

// DeletePhoto handles deleting a photo
func DeletePhoto(c *gin.Context) {
	id := c.Param("id")

	var photo models.Photo
	if err := initializers.DB.First(&photo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	initializers.DB.Delete(&photo)
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
*/
