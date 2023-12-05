package mysql

import (
	"context"

	v1 "github.com/enamespace/tpl/internal/api/v1"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func newUser(db *gorm.DB) *user {
	return &user{
		db: db,
	}
}

func (u *user) Create(ctx context.Context, user *v1.User) error {
	return nil
}

func (u *user) Get(ctx context.Context, username string) (*v1.User, error) {
	return nil, nil
}
