package service

import (
	"check-mate/internal/model"
	"check-mate/internal/store"
)

type MessagesService interface {
	//service methods
	CreateMessage(msg *model.Message) error
	GetMessages(chatId int) ([]*model.Message, error)
	CreateChat(chat *model.Chat)error
	CreateContact(contact *model.Contact) error
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
	err := M.store.CreateMessage(msg.Content, msg.SenderID, msg.ChatID)
	
	return err
}

func (M *messagesService)GetMessages(chatId int) ([]*model.Message, error){
	messages, err := M.store.GetMessages(chatId)
	
	return messages, err
}

func (M *messagesService)CreateChat(chat *model.Chat)error{
	err := M.store.CreateChat(chat.FirstMemberID, chat.SecondMemberID)
	
	return err
}

func (M *messagesService)CreateContact(contact *model.Contact)error{
	err := M.store.CreateContact(contact.Name, contact.Surname, contact.Relationship, contact.UserID)

	return err
}