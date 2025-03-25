package config

//var Cfg = AppConfig{
//	App: App{
//		Name: "webook",
//		Port: 8080,
//	},
//	Jwt: JwtConfig{
//		Key: "c0c5091f15628a94ead9f5b7184d918a",
//	},
//	MySQL: MySQL{
//		DNS: "root:root@tcp(localhost:13306)/webook?charset=utf8mb4&parseTime=True",
//	},
//	Redis: Redis{
//		Addr: "localhost:16379",
//	},
//}

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
