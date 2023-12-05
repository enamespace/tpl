package mysql

import (
	"sync"

	"github.com/enamespace/tpl/internal/store"
	"gorm.io/gorm"
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
