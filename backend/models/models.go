package models

import (
	"gorm.io/gorm"
	// "github.com/google/uuid"
)

type LicenseKey struct {
	gorm.Model
	Key  string `gorm:type:uuid;`
	Hwid string
}

type User struct {
	gorm.Model
	Username string
	Password string
}
