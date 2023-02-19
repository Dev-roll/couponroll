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

func GetUsersFromStoreID(storeID uuid.UUID) []User {
	var users []User
	if err := db.Select(&users, "SELECT * FROM users WHERE id IN (SELECT user_id FROM user_stores WHERE store_id=?)", storeID); err != nil {
		panic(err)
	}
	return users
}

func GetStoresFromUserID(userID uuid.UUID) []Store {
	var stores []Store
	// TODO@notchcoder: Get stores of the authenticated user

	return stores
}

func CreateUserStore(us UserStores) {
	// TODO@notchcoder: Insert user_store record
}

func DeleteUserStore(userID uuid.UUID, us UserStores) {
	// TODO@notchcoder: Delete user_store record
}
