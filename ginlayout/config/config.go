package config

var Cfg AppConfig

type AppConfig struct {
	App   App
	Jwt   JwtConfig
	MySQL MySQL
	Redis Redis
}

type Redis struct {
	Addr   string
	Passwd string
}

type MySQL struct {
	DSN string
}

type JwtConfig struct {
	Key string
}

type App struct {
	Name string
	Port int64
}
