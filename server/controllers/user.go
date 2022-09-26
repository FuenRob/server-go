package controllers

import (
	"fmt"
	"net/http"

	"github.com/fuenrob/server-go/model"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		users := []*model.User{}
		result := db.Order("id asc").Find(&users)

		if result.Error != nil {
			fmt.Print(result.Error)
		}

		return c.JSON(http.StatusOK, users)
	}
}

func GetUserById(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var user model.User

		result := db.First(&user, id)

		if result.Error != nil {
			fmt.Print(result.Error)
		}

		return c.JSON(http.StatusOK, user)
	}
}

func CreateNewUser(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		user := &model.User{Name: name}

		result := db.Create(&user)

		if result.Error != nil {
			fmt.Print(result.Error)
		}

		return c.JSON(http.StatusOK, user)
	}
}

func UpdateUserById(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
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
			return c.JSON(http.StatusInternalServerError, "Error")
		}

		return c.JSON(http.StatusOK, user)
	}
}

func DeleteUserById(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")

		result := db.Delete(&model.User{}, id)

		if result.Error != nil {
			fmt.Print(result.Error)
			return c.JSON(http.StatusInternalServerError, "Error")
		}

		return c.String(http.StatusOK, "Usuario eliminado")
	}
}