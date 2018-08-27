package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Logger Middleware
func Logger() echo.MiddlewareFunc {
	out, err := os.Create("public/logs.txt")
	if err != nil {
		out = os.Stdout
	}

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Method=${method}, Url=\"${uri}\", Status=${status}, Latency:${latency_human} \n",
		Output: out,
	})
}
