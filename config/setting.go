package config

import (
	"github.com/go-ini/ini"
	xerrors "github.com/pkg/errors"
)

var (
	Cfg *ini.File
	err error
)

func init() {
	Cfg, err = ini.Load("./config.ini")
	if err != nil {
		xerrors.Wrapf(err, "Fail to Parse 'config.ini'")
	}
}
