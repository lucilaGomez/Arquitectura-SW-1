package controllers

import (
	"net/http"
	"proyecto/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllAmenities handles fetching all amenities
func GetAllAmenities(c *gin.Context) {
	amenities, err := services.GetAllAmenities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch amenities", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, amenities)
}

// GetAmenityByID handles fetching an amenity by its ID
func GetAmenityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenity ID"})
		return
	}

	amenity, err := services.GetAmenityByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Amenity not found", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, amenity)
}

// CreateAmenity handles creating a new amenity
func CreateAmenity(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body", "details": err.Error()})
		return
	}

	amenity, err := services.CreateAmenity(body.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create amenity", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Amenity created successfully", "amenity": amenity})
}

// UpdateAmenity handles updating an existing amenity
func UpdateAmenity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenity ID"})
		return
	}

	var body struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body", "details": err.Error()})
		return
	}

	amenity, err := services.UpdateAmenity(id, body.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update amenity", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Amenity updated successfully", "amenity": amenity})
}

// DeleteAmenity handles deleting an amenity
func DeleteAmenity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenity ID"})
		return
	}

	if err := services.DeleteAmenity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete amenity", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Amenity deleted successfully"})
}
