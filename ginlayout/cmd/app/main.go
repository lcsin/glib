package main

import (
	"fmt"

	"github.com/lcsin/glib/ginlayout/config"
	"github.com/lcsin/glib/ginlayout/ioc"
	"github.com/lcsin/glib/pkg/ilogrus"
	"github.com/sirupsen/logrus"
)

func main() {
	ioc.InitConfig()
	hook, err := ilogrus.NewErrorHook()
	if err != nil {
		panic(err)
	}
	logrus.AddHook(hook)

	server := InitWebServer()
	server.Run(fmt.Sprintf(":%v", config.Cfg.App.Port))
}
