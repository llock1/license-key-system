package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string

	EmailVerifiedAt *time.Time
	PasswordResetAt *time.Time

	IsSuperAdmin bool
	IsAdmin      bool
	IsModerator  bool
	IsSupport    bool
	IsStaff      bool
	IsBanned     bool
}
