package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserPoints struct {
		UserID    uuid.UUID `json:"user_id,omitempty" db:"user_id"`
		StoreID   uuid.UUID `json:"store_id,omitempty" db:"store_id"`
		Points    int       `json:"points,omitempty" db:"points"`
		CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
		ExpiresAt time.Time `json:"expires_at,omitempty" db:"expires_at"`
	}
)
