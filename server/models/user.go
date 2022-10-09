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
	if hashedPassword, err := HashPassword(user.Password); err == nil {
		db.Statement.SetColumn("Password", hashedPassword)
	}

	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyLogin(db *gorm.DB, username, password string) bool {
	hashedPassword, _ := HashPassword(password)

	var user User
	exist := db.Where("email = ?", username).First(&user)

	if exist.Error != nil {
		return false
	}

	return CheckPasswordHash(hashedPassword, password)
}
