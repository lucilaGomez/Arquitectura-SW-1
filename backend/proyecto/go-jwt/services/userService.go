// services/userService.go

package services

import (
	"errors"
	"net/http"
	"os"
	"proyecto/dtos"
	"proyecto/initializers"
	"proyecto/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body", "details": err.Error()})
		return
	}

	if body.Email == "" || body.Password == "" || body.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	if body.Role != "admin" && body.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password", "details": err.Error()})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash), Role: body.Role}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(dto dtos.LoginUserDTO, c *gin.Context) (*models.User, error) {
	var user models.User
	if err := initializers.DB.First(&user, "email = ?", dto.Email).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, err
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	return &user, nil
}

func Validate(c *gin.Context) (models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return models.User{}, errors.New("Unauthorized")
	}

	return user.(models.User), nil
}
