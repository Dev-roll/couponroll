package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	FavoriteStores struct {
		UserID    uuid.UUID `json:"user_id,omitempty" db:"user_id"`
		StoreID   uuid.UUID `json:"store_id,omitempty" db:"store_id"`
		CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	}
)
