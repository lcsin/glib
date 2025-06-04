package ioc

import (
	"github.com/lcsin/glib/ginlayout/config"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitConfig() {
	fp := pflag.String("config", "config/config.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigFile(*fp)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config.Cfg); err != nil {
		panic(err)
	}
}
