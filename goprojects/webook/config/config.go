package config

var Cfg = AppConfig{
	Jwt: JwtConfig{
		Secret: "c0c5091f15628a94ead9f5b7184d918a",
	},
}

type AppConfig struct {
	Jwt JwtConfig
}

type JwtConfig struct {
	Secret string
}
