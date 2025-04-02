package database

import (
	"fmt"
	"license/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Connect() {
	var err error

	dsn := "postgresql://root:postgres@127.0.0.1:5432/lks"
	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := Client.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		panic(fmt.Sprintf("failed to create extension \"uuid-ossp\", error: %v", err))
	}

	err = Client.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	err = Client.AutoMigrate(&models.Product{})
	if err != nil {
		panic(err)
	}

	err = Client.AutoMigrate(&models.License{})
	if err != nil {
		panic(err)
	}
}
