package cmd

import (
	"fmt"

	"ginlayout/ioc"

	"github.com/spf13/viper"
)

func Run() {
	ioc.InitLocalConfig()
	ioc.InitLogger()
	r := InitWebServer()
	if err := r.Run(fmt.Sprintf(":%v", viper.Get("app.port"))); err != nil {
		panic(err)
	}
}
