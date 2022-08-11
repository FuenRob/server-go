package config

import "fmt"

const (
	DBHost     = "serverdb"
	DBPort     = "3306"
	DBUser     = "root"
	DBPassword = "root"
	DBName     = "mydatabase"
)

func GetConnection() string {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)
	return connection
}
