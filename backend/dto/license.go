package dto

import "time"

type LicenseDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CreatorID int  `json:"creator_id"`
	ProductID int  `json:"product_id"`
	UserID    *int `json:"user_id"`

	Key  string `json:"key"`
	HWID string `json:"hwid"`
}
