package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Coupon struct {
		ID          uuid.UUID `json:"id,omitempty" db:"id"`
		StoreID     uuid.UUID `json:"store_id,omitempty" db:"store_id"`
		Name        string    `json:"name,omitempty" db:"name"`
		Description string    `json:"description,omitempty" db:"description"`
		ImgUrl      string    `json:"img_url,omitempty" db:"img_url"`
		Discount    int       `json:"discount,omitempty" db:"discount"`
		CreatorID   uuid.UUID `json:"creator_id,omitempty" db:"creator_id"`
		CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
		// DeletedAt   time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
		ExpiresAt   time.Time `json:"expires_at,omitempty" db:"expires_at"`
	}
)
