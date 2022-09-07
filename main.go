package main

import (
	"fmt"
	"net/http"

	model "github.com/fuenrob/server-go/model"
	utils "github.com/fuenrob/server-go/utils"
	"github.com/labstack/echo/v4"
)

func main() {

	db := utils.InitConnection()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", func(c echo.Context) error {
		users := []*model.User{}
		result := db.Order("id asc").Find(&users)

		if result.Error != nil {
			fmt.Print(result.Error)
		}

		return c.JSON(http.StatusOK, users)
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		var user model.User
		result := db.First(&user, id)

		if result.Error != nil {
			fmt.Print(result.Error)
		}

		return c.JSON(http.StatusOK, user)
	})

	e.POST("/users", func(c echo.Context) error {
		name := c.FormValue("name")
		user := &model.User{Name: name}

		result := db.Create(&user)

		if result.Error != nil {
			fmt.Print(result.Error)
		}

		return c.JSON(http.StatusOK, user)
	})

	e.PUT("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		name := c.FormValue("name")
		var user model.User
		resultGet := db.First(&user, id)

		if resultGet.Error != nil {
			fmt.Print(resultGet.Error)
		}

		user.Name = name

		resultSave := db.Save(&user)

		if resultSave.Error != nil {
			fmt.Print(resultSave.Error)
		}

		return c.JSON(http.StatusOK, user)

	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id := c.Param("id")

		result := db.Delete(&model.User{}, id)

		if result.Error != nil {
			fmt.Print(result.Error)
		}

		return c.String(http.StatusOK, "¡Usuario Eliminado!")

	})

	e.Logger.Fatal(e.Start(":9999"))
}
