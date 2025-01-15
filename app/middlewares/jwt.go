package middlewares

import (
	"os"
	"strings"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Jwt() echo.MiddlewareFunc {
	secret := os.Getenv("APP_KEY")
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		ContextKey: "token",
		Skipper: middleware.Skipper(func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/login")
		}),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(401, map[string]interface{}{
				"message": "Unauthorized",
			})
		},
	})
}
