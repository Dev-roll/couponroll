package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Store struct {
		ID          uuid.UUID `json:"id,omitempty" db:"id"`
		Name        string    `json:"name,omitempty" db:"name"`
		Description string    `json:"description,omitempty" db:"description"`
		IsPublic    bool      `json:"is_public,omitempty" db:"is_public"`
		IconUrl     string    `json:"icon_url,omitempty" db:"icon_url"`
		ImgUrl      string    `json:"img_url,omitempty" db:"img_url"`
		CreatorID   uuid.UUID `json:"creator_id,omitempty" db:"creator_id"`
		CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
		// DeletedAt   time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	}
)

func CreateStore(s Store) {
	v, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("INSERT INTO stores (id, name, description, is_public, icon_url, img_url, creator_id) VALUES (?, ?, ?, ?, ?, ?, ?)", v, s.Name, s.Description, s.IsPublic, s.IconUrl, s.ImgUrl, s.CreatorID); err != nil {
		panic(err)
	}
}

func GetStoreFromID(s Store) Store {
	var store Store
	if err := db.Get(&store, "SELECT * FROM stores WHERE id=?", s.ID); err != nil {
		panic(err)
	}
	return store
}

func UpdateStore(storeID uuid.UUID, s Store) {
	if _, err := db.Exec("UPDATE stores SET name=?, description=?, is_public=?, icon_url=?, img_url=? WHERE id=?", s.Name, s.Description, s.IsPublic, s.IconUrl, s.ImgUrl, storeID); err != nil {
		panic(err)
	}
}

func DeleteStore(s Store) {
	if _, err := db.Exec("DELETE FROM stores WHERE id=?", s.ID); err != nil {
		panic(err)
	}
}
