package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID        uuid.UUID `json:"id,omitempty" db:"id"`
		Name      string    `json:"name,omitempty" db:"name"`
		Email     string    `json:"email,omitempty" db:"email"`
		Password  string    `json:"password,omitempty" db:"password"`
		Role      string    `json:"role,omitempty" db:"role"`
		CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
	}
)

func GetUserFromID(u User) User {
	var user User
	if err := db.Get(&user, "SELECT * FROM users WHERE id=?", u.ID); err != nil {
		// TODO: handle error
		panic(err)
	}
	return user
}

func GetUserFromName(u User) User {
	var user User
	if err := db.Get(&user, "SELECT * FROM users WHERE name=?", u.Name); err != nil {
		// TODO: handle error
		panic(err)
	}
	return user
}

func GetUserFromEmail(u User) User {
	var user User
	if err := db.Get(&user, "SELECT * FROM users WHERE email=?", u.Email); err != nil {
		// TODO: handle error
		fmt.Println(err)
		panic(err)
	}
	return user
}

func CreateUser(u User) {
	v, err := uuid.NewRandom()
	if err != nil {
		// TODO: handle error
		panic(err)
	}

	if _, err := db.Exec("INSERT INTO users (id, name, email, password, role) VALUES (?, ?, ?, ?, ?)", v, u.Name, u.Email, u.Password, u.Role); err != nil {
		// TODO: handle error
		panic(err)
	}
}

func GetUsersCountFromName(u User) int {
	var count int
	if err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE name=?", u.Name); err != nil {
		// TODO: handle error
		panic(err)
	}
	return count
}

func UpdateUser(userID uuid.UUID, u User) {
	if _, err := db.Exec("UPDATE users SET name=?, email=?, role=? WHERE id=?", u.Name, u.Email, u.Role, userID); err != nil {
		// TODO: handle error
		panic(err)
	}
}

func UpdatePassword(userID uuid.UUID, u User) {
	if _, err := db.Exec("UPDATE users SET password=? WHERE id=?", u.Password, userID); err != nil {
		panic(err)
	}
}
