package helpers

import (
	"goBoilterplate/app/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// AuthMakeToken helper to make a token from user
func AuthMakeToken(user *models.User) (string, error) {
	secret := os.Getenv("APP_KEY")
	claims := models.JwtClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

// AuthGetUser helper to retrieve the user from token
func AuthGetUser(c echo.Context) *models.User {
	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(*models.JwtClaims)

	user := models.UserShow(claims.ID)
	if user != nil {
		return user
	}
	return nil
}
