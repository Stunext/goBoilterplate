package middlewares

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Gzip Middleware
func Gzip() echo.MiddlewareFunc {
	return middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Path(), "docs") {
				return true
			}
			return false
		},
	})
}
