package models

import (
	"gorm.io/gorm"
)

type License struct {
	gorm.Model

	CreatorID int
	Creator   User
	ProductID int
	Product   Product
	UserID    *int
	User      *User

	Key string `gorm:"uniqueIndex"`
}
