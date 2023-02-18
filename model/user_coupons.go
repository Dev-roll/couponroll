package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserCoupons struct {
		UserID    uuid.UUID `json:"user_id,omitempty" db:"user_id"`
		CouponID  uuid.UUID `json:"coupon_id,omitempty" db:"coupon_id"`
		IsUsed    bool      `json:"is_used,omitempty" db:"is_used"`
		CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
	}
)
