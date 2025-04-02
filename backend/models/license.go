package models

import (
	"gorm.io/gorm"
	// "github.com/google/uuid"
)

type License struct {
	gorm.Model

	Creator User
	Product Product

	Key  string `gorm:"type:uuid"`
	Hwid string
}
