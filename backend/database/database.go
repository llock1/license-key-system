package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"license/models"
)

var Client *gorm.DB

func Connect() {
	var err error
	Client, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Client.AutoMigrate(&models.License{})
	Client.AutoMigrate(&models.User{})
}
