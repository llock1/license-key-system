package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"license/models"
)

var Db *gorm.DB
var err error

func Connect() {
	Db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&models.LicenseKey{})
}