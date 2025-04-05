package dto

import "time"

type UserDTO struct {
	Username string  `json:"username"`
	Email    *string `json:"email"`

	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	PasswordResetAt *time.Time `json:"password_reset_at"`

	IsSuperAdmin bool `json:"is_super_admin"`
	IsAdmin      bool `json:"is_admin"`
	IsModerator  bool `json:"is_moderator"`
	IsSupport    bool `json:"is_support"`
	IsStaff      bool `json:"is_staff"`
	IsBanned     bool `json:"is_banned"`
}
