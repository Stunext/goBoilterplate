package main

import (
	"goBoilterplate/app/console"
	"goBoilterplate/app/router"
	"goBoilterplate/config"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"
)

// @title Golang Echo API
// @version 1.0
// @description API documentation by Swagger

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	app := echo.New()

	db := config.Database()
	defer db.Close()

	config.Redis()
	console.Schedule()
	router.Init(app)

	app.Server.Addr = ":3000"
	graceful.ListenAndServe(app.Server, 5*time.Second)
}
