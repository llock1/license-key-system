package dto

import "time"

type SubscriptionDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID      int        `json:"user_id"`
	ProductID   int        `json:"product_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	HWID        *string    `json:"hwid"`
	HWIDResetAt *time.Time `json:"hwid_reset_at"`

	IsRevoked bool `json:"is_revoked"`
}
