package main

import (
	"ClassWork10-23/config"
	"github.com/pkg/errors"
)

func manage() (err error) {
	defer func() {err = errors.Wrap(err, "main.manage")}()

	conf = new(config.Config)

	err = config.LoadConfigHCL("./config.hcl", conf)
	if err != nil {
		err = errors.Wrap(err, "ошибка загрузки конфигурации")
	}
	err = dbase.Connection(conf.Data.DataPass, conf.Data.UserName, conf.Data.Password)
	if err != nil {
		return
	}

	return
}