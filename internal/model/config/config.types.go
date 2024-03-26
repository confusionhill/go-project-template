package config

type Config struct {
	App AppConfig      `json:"app"`
	Sql DatabaseConfig `json:"sql"`
	Jwt JWTConfig      `json:"jwt"`
}

type AppConfig struct {
	Name string `json:"name"`

	//Log level, 1 for debug 2 for production
	Level int64  `json:"level"`
	Port  string `json:"port"`
}

type DatabaseConfig struct {
	Host string `json:"host"`
	Type string `json:"type"`
}

type PaymentConfig struct {
	ApiKey string `json:"api_key"`
	Host   string `json:"host"`
}

type JWTConfig struct {
	Secret string `json:"secret"`
	Span   int64  `json:"span"`
}
