package controllers

import (
	"net/http"

	model "github.com/fuenrob/server-go/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		if !model.VerifyLogin(db, email, password) {
			return echo.ErrUnauthorized
		}

		token := model.GenerateUserToken(email)

		return c.JSON(http.StatusOK, echo.Map{
			"error": false,
			"token": token,
		})
	}
}
