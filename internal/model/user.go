package model

import "time"

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string
	CreatedAt time.Time
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	ID        string `db:"id" json:"token"`
	UserID    string `db:"user_id"`
	CreatedAt string `db:"created_at"`
}

type Note struct {
}
