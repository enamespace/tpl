package mysql

import (
	"fmt"
	"sync"

	"gorm.io/gorm"

	"github.com/enamespace/tpl/internal/store"
	mysqldb "github.com/enamespace/tpl/pkg/db/mysql"
	genericoptions "github.com/enamespace/tpl/pkg/options"
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

func GetMysqlFactory(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		dbIns, err = mysqldb.New(opts, nil)

		mysqlFactory = &datastore{
			db: dbIns,
		}

	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}
