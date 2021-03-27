package Model

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	xerrors "github.com/pkg/errors"
)

var MyDB *gorm.DB

func Init(cfg *ini.File) error {
	sec, err := cfg.GetSection("database")
	if err != nil {
		return xerrors.Wrapf(err, "Fail to get section 'database'")
	}
	dbType := sec.Key("TYPE").String()
	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()

	MyDB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		return xerrors.Wrapf(err, "Open db fail")
	}

	MyDB.SingularTable(true)
	MyDB.LogMode(true)
	MyDB.DB().SetMaxIdleConns(10)
	MyDB.DB().SetMaxOpenConns(100)

	MyDB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	MyDB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	if !MyDB.HasTable(&UserModel{}) {
		MyDB.AutoMigrate(&UserModel{})
	}

	return nil
}
