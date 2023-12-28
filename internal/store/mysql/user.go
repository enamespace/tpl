package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	v1 "github.com/enamespace/tpl/internal/api/v1"
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
	return u.db.Create(user).Error
}

func (u *user) Update(ctx context.Context, user *v1.User) error {
	return u.db.Save(user).Error
}

func (u *user) Get(ctx context.Context, username string) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("name = ? and status = 1", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// TODO: fake delete
func (u *user) Delete(ctx context.Context, username string) error {

	err := u.db.Where("name = ?", username).Delete(&v1.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("database err when delete user")
	}
	return err
}
