package store

import (
	"context"

	v1 "github.com/enamespace/tpl/internal/api/v1"
)

type UserStore interface {
	Create(ctx context.Context, user *v1.User) error
	Get(ctx context.Context, username string) (*v1.User, error)
	Update(ctx context.Context, user *v1.User) error
	Delete(ctx context.Context, username string) error
}
