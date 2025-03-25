package cmd

import (
	"fmt"

	"github.com/lcsin/webook/ioc"
	"github.com/spf13/viper"
)

func Run() {
	ioc.InitConfig()
	server := InitWebServer()
	server.Run(fmt.Sprintf(":%v", viper.Get("app.port")))
}
