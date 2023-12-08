package service

import (
	"check-mate/internal/model"
	"check-mate/internal/store"
)

type UsersService interface {
	CreateUser(user *model.User) (uint, error)
	GetUser(username string) (*model.User, error)
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

func (s *service) CreateUser(user *model.User) (uint, error) {
	id, err := s.store.CreateUser(user)

	return id, err
}

func (s *service) GetUser(username string) (*model.User, error) {
	user, err := s.store.GetUser(username)

	return user, err
}
