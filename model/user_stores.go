package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserStores struct {
		UserID   uuid.UUID `json:"user_id,omitempty" db:"user_id"`
		StoreID  uuid.UUID `json:"store_id,omitempty" db:"store_id"`
		JoinedAt time.Time `json:"joined_at,omitempty" db:"joined_at"`
	}
)
