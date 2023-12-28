package mysql

import (
	"sync"

	"gorm.io/gorm"

	"github.com/enamespace/tpl/internal/store"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) User() store.UserStore {
	return newUser(ds.db)
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMysqlFactory() (store.Factory, error) {
	once.Do(func() {
		mysqlFactory = &datastore{}
	})

	return mysqlFactory, nil
}
