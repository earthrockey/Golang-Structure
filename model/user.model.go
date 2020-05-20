package model

import "time"

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Username  string `json:"username"`
	Password  string `json:"password"`
}
