package repository

import (
	. "UserService/Domain/model"
	"context"
	"fmt"
	"github.com/go-ini/ini"
	xerrors "github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func Init(cfg *ini.File) error {
	sec, err := cfg.GetSection("database")
	if err != nil {
		return xerrors.Wrapf(err, "Fail to get section 'database'")
	}

	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return xerrors.Wrapf(err, "Open DB Fail")
	}

	MyDB, err := db.DB()
	if err != nil {
		return xerrors.Wrapf(err, "")
	}
	MyDB.SetMaxIdleConns(10)
	MyDB.SetMaxOpenConns(100)

	if !db.Migrator().HasTable(UserModel{}) {
		err := db.AutoMigrate(UserModel{})
		if err != nil {
			return xerrors.Wrapf(err, "Create Table Fail")
		}
	}

	GlobalDB = db

	return nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return GlobalDB.WithContext(ctx)
}
