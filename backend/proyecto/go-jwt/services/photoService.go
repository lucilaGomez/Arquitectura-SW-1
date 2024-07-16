package services

import (
	"errors"
	"proyecto/initializers"
	"proyecto/models"
)

// CreatePhoto crea una nueva foto en la base de datos.
func CreatePhoto(photo *models.Photo) error {
	if err := initializers.DB.Create(&photo).Error; err != nil {
		return err
	}
	return nil
}

// GetAllPhotos obtiene todas las fotos desde la base de datos.
func GetAllPhotos() ([]models.Photo, error) {
	var photos []models.Photo
	if err := initializers.DB.Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

// GetPhotoByID obtiene una foto por su ID desde la base de datos.
func GetPhotoByID(id string) (models.Photo, error) {
	var photo models.Photo
	if err := initializers.DB.First(&photo, id).Error; err != nil {
		return models.Photo{}, errors.New("photo not found")
	}
	return photo, nil
}

// UpdatePhoto actualiza los detalles de una foto existente en la base de datos.
func UpdatePhoto(id string, updatedPhoto *models.Photo) error {
	var existingPhoto models.Photo
	if err := initializers.DB.First(&existingPhoto, id).Error; err != nil {
		return errors.New("photo not found")
	}

	// Actualizar solo los campos modificados
	existingPhoto.URL = updatedPhoto.URL
	existingPhoto.HotelID = updatedPhoto.HotelID

	if err := initializers.DB.Save(&existingPhoto).Error; err != nil {
		return err
	}
	return nil
}

// DeletePhoto elimina una foto por su ID desde la base de datos.
func DeletePhoto(id string) error {
	var photo models.Photo
	if err := initializers.DB.First(&photo, id).Error; err != nil {
		return errors.New("photo not found")
	}

	if err := initializers.DB.Delete(&photo).Error; err != nil {
		return err
	}
	return nil
}
