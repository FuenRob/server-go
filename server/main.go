package main

import (
	controllers "github.com/fuenrob/server-go/controllers"
	utils "github.com/fuenrob/server-go/utils"
	"github.com/labstack/echo"
)

func main() {

	db := utils.InitConnection()

	e := echo.New()

	e.GET("/users", controllers.GetAllUsers(db))
	e.GET("/users/:id", controllers.GetUserById(db))
	e.POST("/users", controllers.CreateNewUser(db))
	e.PUT("/users/:id", controllers.UpdateUserById(db))
	e.DELETE("/users/:id", controllers.DeleteUserById(db))

	e.Logger.Fatal(e.Start(":9999"))
}
