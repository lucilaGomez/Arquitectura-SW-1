package services

import (
	"proyecto/initializers"
	"proyecto/models"
)

// GetAllAmenities returns all amenities
func GetAllAmenities() ([]models.Amenity, error) {
	var amenities []models.Amenity
	if err := initializers.DB.Find(&amenities).Error; err != nil {
		return nil, err
	}
	return amenities, nil
}

// GetAmenityByID returns an amenity by its ID
func GetAmenityByID(id int) (models.Amenity, error) {
	var amenity models.Amenity
	if err := initializers.DB.First(&amenity, id).Error; err != nil {
		return models.Amenity{}, err
	}
	return amenity, nil
}

// CreateAmenity creates a new amenity
func CreateAmenity(name string) (models.Amenity, error) {
	amenity := models.Amenity{Name: name}
	if err := initializers.DB.Create(&amenity).Error; err != nil {
		return models.Amenity{}, err
	}
	return amenity, nil
}

// UpdateAmenity updates an existing amenity
func UpdateAmenity(id int, name string) (models.Amenity, error) {
	var amenity models.Amenity
	if err := initializers.DB.First(&amenity, id).Error; err != nil {
		return models.Amenity{}, err
	}
	amenity.Name = name
	if err := initializers.DB.Save(&amenity).Error; err != nil {
		return models.Amenity{}, err
	}
	return amenity, nil
}

// DeleteAmenity deletes an amenity by its ID
func DeleteAmenity(id int) error {
	if err := initializers.DB.Delete(&models.Amenity{}, id).Error; err != nil {
		return err
	}
	return nil
}
