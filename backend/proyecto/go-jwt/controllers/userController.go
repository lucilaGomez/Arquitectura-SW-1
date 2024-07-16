// controllers/userController.go

package controllers

import (
	"net/http"
	"proyecto/dtos"
	"proyecto/initializers"
	"proyecto/models"
	"proyecto/services"

	"github.com/gin-gonic/gin"
)

// SignUp controller function
func SignUp(c *gin.Context) {
	services.SignUp(c)
}

// Login controller function
func Login(c *gin.Context) {
	var dto dtos.LoginUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := services.Login(dto, c) // Pasar dto por valor en lugar de referencia
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Validate controller function
func Validate(c *gin.Context) {
	user, err := services.Validate(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var foundUser models.User
	if err := initializers.DB.First(&foundUser, user.(models.User).ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": foundUser})
}

func Logout(c *gin.Context) {
	// Eliminar la cookie de autorización
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
