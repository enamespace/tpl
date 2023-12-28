package service

import (
	"context"

	v1 "github.com/enamespace/tpl/internal/api/v1"
	"github.com/enamespace/tpl/internal/store"
)

type UserService interface {
	Create(ctx context.Context, User *v1.User) error
	Update(ctx context.Context, User *v1.User) error
	Get(ctx context.Context, username string) (*v1.User, error)
	Delete(ctx context.Context, username string) error
}

type userService struct {
	ds store.Factory
}

func newUserSerivce(ds store.Factory) *userService {
	return &userService{
		ds: ds,
	}
}

func (u *userService) Create(ctx context.Context, user *v1.User) error {
	return u.ds.User().Create(ctx, user)
}

func (u *userService) Update(ctx context.Context, user *v1.User) error {
	return u.ds.User().Create(ctx, user)
}

func (u *userService) Get(ctx context.Context, username string) (*v1.User, error) {
	return u.ds.User().Get(ctx, username)
}

func (u *userService) Delete(ctx context.Context, username string) error {
	return u.ds.User().Delete(ctx, username)
}
