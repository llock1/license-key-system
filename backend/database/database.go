package database

import (
	"license/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Connect() {
	Db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&models.LicenseKey{})
}
