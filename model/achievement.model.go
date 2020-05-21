package model

import "github.com/jinzhu/gorm"

type AchievementRequest struct {
	Name string `json:"name"`
}

type Achievement struct {
	gorm.Model
	Name string `json:"name"`
}