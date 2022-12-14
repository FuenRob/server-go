package utils

import (
	"log"

	config "github.com/fuenrob/server-go/configs"
	model "github.com/fuenrob/server-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnection() *gorm.DB {
	dsn := config.GetConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Connection error")
	}
	db.AutoMigrate(&model.User{})

	return db
}
