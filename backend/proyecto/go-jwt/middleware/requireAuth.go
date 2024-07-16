package middleware

import (
	"fmt"
	"net/http"
	"os"
	"proyecto/initializers"
	"proyecto/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Obtener la cookie de la solicitud
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decodificar y validar el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Verificar la expiración del token
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Buscar el usuario con el ID del token
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Adjuntar el usuario a la solicitud
		c.Set("user", user)

		// Continuar
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func RequireAdmin(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if user.(models.User).Role != "admin" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Next()
}
