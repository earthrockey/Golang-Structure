package model

import (
	"github.com/jinzhu/gorm"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Username    string        `json:"username"`
	Password    string        `json:"password"`
	Book        []Book        `json:"book"`
	Achievement []Achievement `json:"achievement" gorm:"many2many:user_achievements;"`
}
