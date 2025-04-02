package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Owner       User
	Title       string
	Description string

	IsListed bool
}
