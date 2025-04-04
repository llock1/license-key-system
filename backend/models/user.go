package models

import (
	"license/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
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

func (u *User) SetPassword(password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
