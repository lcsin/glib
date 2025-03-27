package cmd

import (
	"fmt"

	"github.com/lcsin/webook/config"
	"github.com/lcsin/webook/ioc"
)

func Run() {
	ioc.InitConfig()
	server := InitWebServer()
	server.Run(fmt.Sprintf(":%v", config.Cfg.App.Port))
}
