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

const (
	RoleUser = iota
	RoleStaff
	RoleSupport
	RoleModerator
	RoleAdmin
	RoleSuperAdmin
)

func (user *User) GetUserRole() int {
	switch {
	case user.IsSuperAdmin:
		return RoleSuperAdmin
	case user.IsAdmin:
		return RoleAdmin
	case user.IsModerator:
		return RoleModerator
	case user.IsSupport:
		return RoleSupport
	case user.IsStaff:
		return RoleStaff
	default:
		return RoleUser
	}
}

func (user *User) IsAbove(other *User) bool {
	return user.GetUserRole() < other.GetUserRole()
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
