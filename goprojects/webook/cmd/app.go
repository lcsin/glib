package cmd

import (
	"fmt"

	"github.com/lcsin/webook/config"
)

func Run() {
	//server := InitWebServer()
	//server.Run(config.Cfg)

	server := InitWebServer()
	server.Run(fmt.Sprintf(":%v", config.Cfg.App.Port))
}
