package options

import "time"

// MySQLOptions defines options for mysql database
type MySQLOptions struct {
	Host                  string        `json:"host,omitempty"                      mapstructure:"host"`
	Username              string        `json:"username,omitempty"                  mapstructure:"username"`
	Password              string        `json:"-"                                   mapstructure:"password"`
	Database              string        `json:"database"                            mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty"      mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty"      mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty"  mapstructure:"max-connection-life-time"`
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Host:                  "127.0.0.1:3306",
		Username:              "",
		Password:              "",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

func (o *MySQLOptions) Validate() []error {
	errs := []error{}
	return errs
}
