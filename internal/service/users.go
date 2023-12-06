package service

import "check-mate/internal/store"

type UsersService interface {
	//service methods
}

// service struct that implements Users store above
type service struct {
	store store.Store
}

func NewUserService(s store.Store) UsersService {
	return &service{
		store: s,
	}
}
