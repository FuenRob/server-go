package controllers

import (
	"net/http"

	model "github.com/fuenrob/server-go/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := []byte(c.FormValue("password"))
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

		if err != nil {
			return c.JSON(http.StatusConflict, echo.Map{
				"error":        true,
				"token":        "",
				"errorMessage": "Problema al encriptar la contraseña",
			})
		}

		var user model.User
		exist := db.Where("email = ?", email).First(&user)

		if exist.Error != nil {
			return echo.ErrUnauthorized
		}

		errMatch := bcrypt.CompareHashAndPassword(hashedPassword, password)

		if errMatch != nil {
			return echo.ErrUnauthorized
		}

		token := model.GenerateUserToken(email)

		return c.JSON(http.StatusOK, echo.Map{
			"error": false,
			"token": token,
		})
	}
}
