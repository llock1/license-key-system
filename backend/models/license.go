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

	Key  string `gorm:"uniqueIndex"`
	Hwid string
}
