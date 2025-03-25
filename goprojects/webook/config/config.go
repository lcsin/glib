package config

var Cfg = AppConfig{
	App: App{
		Port: "8080",
	},
	Jwt: JwtConfig{
		Secret: "c0c5091f15628a94ead9f5b7184d918a",
	},
	MySQL: MySQL{
		DNS: "root:root@tcp(localhost:13306)/webook?charset=utf8mb4&parseTime=True",
	},
	Redis: Redis{
		Addr: "localhost:16379",
	},
}

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
	DNS string
}

type JwtConfig struct {
	Secret string
}

type App struct {
	Port string
}
