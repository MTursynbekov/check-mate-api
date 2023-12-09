package service

import (
	"check-mate/internal/model"
	"check-mate/internal/store"
)

type ContactsService interface {
	CreateContact(contact *model.Contact) error
	GetContacts(userId int) ([]*model.Contact, error)
	GetContact(userId, chatId int) (*model.Contact, error)
}

type contactService struct {
	store store.Store
}

func NewContactService(s store.Store) ContactsService {
	return &contactService{
		store: s,
	}
}

func (s *contactService) CreateContact(contact *model.Contact) error {
	err := s.store.CreateContact(contact)

	return err
}

func (s *contactService) GetContacts(userId int) ([]*model.Contact, error) {
	contacts, err := s.store.GetContacts(userId)

	return contacts, err
}

func (s *contactService) GetContact(userId, chatId int) (*model.Contact, error) {
	contact, err := s.store.GetContact(userId, chatId)

	return contact, err
}
