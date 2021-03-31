package model

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UserName string `gorm:"not null" gorm:"unique" gorm:"size:10" json:"user_name"`
	Password string `gorm:"not null" json:"password"`
	Member   bool   `json:"member"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Like     int    `json:"like"`
	About    string `gorm:"size:120" json:"about"`
}
