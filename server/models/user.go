package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (user *User) BeforeSave(db *gorm.DB) (err error) {
	password := []byte(user.Password)
	if hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost); err == nil {
		db.Statement.SetColumn("Password", hashedPassword)
	}

	return err
}
