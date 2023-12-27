package mysql

import (
	"fmt"

	"github.com/enamespace/tpl/pkg/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient(o *options.MySQLOptions, opts ...gorm.Option) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local",
	)

	db, err := gorm.Open(mysql.Open(dsn), opts...)

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(o.MaxConnectionLifeTime)

	return db, nil
}
