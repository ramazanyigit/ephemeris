package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func CreateConnection() {
	var err error
	dsn := "root:ephemeris@tcp(127.0.0.1:3306)/ephemeris?charset=utf8mb4&parseTime=True&loc=Local"
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetConnection() *gorm.DB {
	return database
}
