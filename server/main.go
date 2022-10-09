package main

import (
	controllers "github.com/fuenrob/server-go/controllers"
	model "github.com/fuenrob/server-go/models"
	utils "github.com/fuenrob/server-go/utils"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func main() {

	db := utils.InitConnection()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", controllers.Login(db))

	r := e.Group("/admin")

	// Configure middleware with the custom claims type
	r.Use(middleware.JWTWithConfig(model.GetConfigJWT()))

	r.GET("/users", controllers.GetAllUsers(db))
	r.GET("/users/:id", controllers.GetUserById(db))
	e.POST("/users", controllers.CreateNewUser(db))
	r.PUT("/users/:id", controllers.UpdateUserById(db))
	r.DELETE("/users/:id", controllers.DeleteUserById(db))

	e.Logger.Fatal(e.Start(":9999"))
}
