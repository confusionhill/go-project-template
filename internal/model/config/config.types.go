package config

type Config struct {
	App AppConfig      `json:"app"`
	Sql DatabaseConfig `json:"sql"`
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
