package cmd

import "ginlayout/ioc"

func Run() {
	ioc.InitLocalConfig()
	ioc.InitLogger()
}
