package helpers

import (
	"goBoilterplate/app/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// GetAuthorizationToken helper to get token from header
func GetAuthorizationToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Unauthorized")
	}

	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), nil
	}

	return "", echo.NewHTTPError(http.StatusBadRequest, "Unauthorized")
}

// AuthMakeToken helper to make a token from user
func AuthMakeToken(user *models.User) (string, error) {
	secret := os.Getenv("APP_KEY")
	claims := models.JwtClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

// AuthGetUser helper to retrieve the user from token
func AuthGetUser(c echo.Context) *models.User {
	secret := os.Getenv("APP_KEY")

	tokenString, err := GetAuthorizationToken(c)
	if err != nil {
		return nil
	}

	token, _ := jwt.ParseWithClaims(tokenString, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*models.JwtClaims); ok {
		user := models.UserShow(claims.ID)
		if user != nil {
			return user
		}
	}
	return nil
}
