package controllers

import (
	"fmt"
	"net/http"

	model "github.com/fuenrob/server-go/models"
	echo "github.com/labstack/echo/v4"
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
		email := c.FormValue("email")
		password := c.FormValue("password")

		user := &model.User{Name: name, Email: email, Password: password}

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
		email := c.FormValue("email")
		password := c.FormValue("password")

		var user model.User

		resultGet := db.First(&user, id)

		if resultGet.Error != nil {
			fmt.Print(resultGet.Error)
		}

		user.Name = name
		user.Email = email
		user.Password = password

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
