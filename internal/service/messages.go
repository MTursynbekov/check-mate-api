package service

import (
	"check-mate/internal/model"
	"check-mate/internal/store"
)

type MessagesService interface {
	//service methods
	CreateMessage(msg *model.Message) error
}

// service struct that implements Users store above
type messagesService struct {
	store store.Store
}

func NewMessagesService(s store.Store) MessagesService {
	return &messagesService{
		store: s,
	}
}

func (M *messagesService)CreateMessage(msg *model.Message)error{
	err := M.store.CreateMessage(msg.Text, msg.SenderID, msg.ReceiverID)
	
	return err
}