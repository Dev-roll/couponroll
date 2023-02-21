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
	if err := db.Select(&stores, "SELECT stores.* FROM stores, user_stores WHERE stores.id = store_id AND user_id=?", userID); err != nil {
		panic(err)
	}
	return stores
}

func CreateUserStore(us UserStores) {
	if _, err := db.Exec("INSERT INTO user_stores (user_id, store_id) VALUES (?, ?)", us.UserID, us.StoreID); err != nil {
		panic(err)
	}
}

func DeleteUserStore(userID uuid.UUID, us UserStores) {
	if _, err := db.Exec("DELETE FROM user_stores WHERE user_id=? AND store_id=?", userID, us.StoreID); err != nil {
		panic(err)
	}
}
