package config

type Config struct {
	LogLevel string
	Port     string
	MySQL    struct {
		Host     string
		Username string
		Password string
		Database string
	}
}
