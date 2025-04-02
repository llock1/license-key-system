package models

import (
	"gorm.io/gorm"
	"time"
	// "github.com/google/uuid"
)

type LicenseKey struct {
	gorm.Model
	Key       string `gorm:"type:uuid"`
	Hwid      string
	ExpiresAt *time.Time
}
