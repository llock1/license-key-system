package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model

	UserID    int
	User      User
	ProductID int
	Product   Product

	StartDate time.Time
	EndDate   *time.Time

	HWID        *string
	HWIDResetAt *time.Time

	IsRevoked bool
}
