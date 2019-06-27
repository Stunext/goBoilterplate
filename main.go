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

// @title Golang Boilterplate
// @version 1.0
// @description This is a sample server

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
