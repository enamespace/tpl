package config

import "github.com/enamespace/tpl/pkg/options"

type Config struct {
	GenericServerOptions options.GenericServerOptions `json:"server"   mapstructure:"server"`
	MySQLOptions         options.MySQLOptions         `json:"mysql"   mapstructure:"mysql"`
}
