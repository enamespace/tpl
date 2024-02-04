package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	v1 "github.com/enamespace/tpl/internal/api/v1"
)

type DBOptions struct {
	host     string
	username string
	password string
	database string

	admin bool
}

func NewDBOptions() *DBOptions {
	return &DBOptions{
		host:     "127.0.0.1:3306",
		username: "root",
		password: "root",
		database: "tpl",
		admin:    false,
	}
}

func (o *DBOptions) Run() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		o.username, o.password, o.host, o.database, true, "Local")

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	if db.Migrator().HasTable(&v1.User{}) {
		err = db.AutoMigrate(&v1.User{})
	} else {
		err = db.Debug().Migrator().CreateTable(&v1.User{})
	}

	return err
}
