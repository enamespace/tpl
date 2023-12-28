package config

import "github.com/enamespace/tpl/internal/options"

type Config struct {
	*options.Options
}

func CreateConfigFromOptions(o *options.Options) *Config {
	return &Config{o}
}
