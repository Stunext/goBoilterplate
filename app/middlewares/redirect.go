package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// HTTPSRedirect Middleware
func HTTPSRedirect() echo.MiddlewareFunc {
	return middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
		Code: http.StatusMovedPermanently,
	})
}

// NonWWWRedirect Middleware
func NonWWWRedirect() echo.MiddlewareFunc {
	return middleware.NonWWWRedirectWithConfig(middleware.RedirectConfig{
		Code: http.StatusMovedPermanently,
	})
}
