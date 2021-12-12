package database

import "github.com/ramazanyigit/ephemeris/internal/model"

func AutoMigration() {
	err := database.AutoMigrate(&model.DiaryLog{})
	if err != nil {
		panic("Cannot migrate DiaryLog: " + err.Error())
	}
}