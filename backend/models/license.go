package models

import (
	"gorm.io/gorm"
	// "github.com/google/uuid"
)

type License struct {
	gorm.Model

	CreatorID int
	Creator   User
	ProductID int
	Product   Product

	Key  string
	Hwid string
}
