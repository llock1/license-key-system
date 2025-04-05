package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string `gorm:"uniqueIndex"`
	OwnerID     int
	Owner       User
	Title       string
	Description string
	Status      string
	Version     string

	IsListed bool
}
