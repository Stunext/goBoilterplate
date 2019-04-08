package router

import (
	"goBoilterplate/app/controllers"
	"goBoilterplate/app/middlewares"
	_ "goBoilterplate/docs" // For Swagger
	
	"log"
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
)

// Init Router
func Init(app *echo.Echo) {
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())

	app.GET("/", controllers.Index)
	app.GET("/test", controllers.Test)
	app.GET("/docs/*", echoSwagger.WrapHandler)

	api := app.Group("/api", middlewares.Jwt())
	{
		api.POST("/login", controllers.Login)
		api.GET("/logout", controllers.Logout)

		users := api.Group("/users")
		{
			users.GET("", controllers.UserList)
			users.POST("", controllers.UserStore)
			users.GET("/:id", controllers.UserShow)
			users.PUT("/:id", controllers.UserUpdate)
			users.DELETE("/:id", controllers.UserDelete)
		}
	}

	log.Printf("Server started...")
}
