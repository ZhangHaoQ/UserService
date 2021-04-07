package config

import (
	"github.com/go-ini/ini"
	xerrors "github.com/pkg/errors"
)

func Init() (*ini.File, error) {
	Cfg, err := ini.Load("D:\\GoPro\\UserService\\config\\config.ini")
	if err != nil {
		return nil, xerrors.Wrapf(err, "Fail to Parse 'config.ini'")
	}
	return Cfg, nil
}
