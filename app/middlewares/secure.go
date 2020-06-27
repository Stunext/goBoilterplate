package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Secure Middleware
func Secure() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            0,
		ContentSecurityPolicy: "",
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Path(), "/docs") {
				return true
			}
			return false
		},
	})
}
