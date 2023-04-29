package models

import (
	"database/sql"
	"time"
)

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

type Models struct {
	User  User
	Token Token
}

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Token struct {
	UserID       int    `json:"user_id" gorm:"primaryKey"'`
	RefreshToken string `json:"refresh_token"`
}
