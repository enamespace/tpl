package service

import "github.com/enamespace/tpl/internal/store"

type Service interface {
	User() UserService
}

type service struct {
	ds store.Factory
}

func NewService(ds store.Factory) *service {
	return &service{
		ds: ds,
	}
}

func (s *service) User() UserService {
	return newUserSerivce(s.ds)
}
