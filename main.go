package main

import (
	"goBoilterplate/app/console"
	"goBoilterplate/app/router"
	"goBoilterplate/config"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
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

	app.Logger.Fatal(app.Start(":3000"))
}
