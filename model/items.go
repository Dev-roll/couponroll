package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Items struct {
		ID          uuid.UUID `json:"id,omitempty" db:"id"`
		StoreID     uuid.UUID `json:"store_id,omitempty" db:"store_id"`
		Name        string    `json:"name,omitempty" db:"name"`
		Description string    `json:"description,omitempty" db:"description"`
		Price       int       `json:"price,omitempty" db:"price"`
		Stock       int       `json:"stock,omitempty" db:"stock"`
		IsPublic    bool      `json:"is_public,omitempty" db:"is_public"`
		CreatorID   uuid.UUID `json:"creator_id,omitempty" db:"creator_id"`
		CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
		// DeletedAt   time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	}
)
