package Model

import (
	"UserService/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init() {
	sec, err := config.Cfg.GetSection("database")
}
