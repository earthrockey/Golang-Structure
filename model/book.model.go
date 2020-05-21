package model

import "github.com/jinzhu/gorm"

type BookRequest struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}
