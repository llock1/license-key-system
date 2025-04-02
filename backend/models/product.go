package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string
	OwnerID     int
	Owner       User
	Title       string
	Description string

	IsListed bool
}
