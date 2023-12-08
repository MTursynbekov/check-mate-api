package store

import (
	"check-mate/internal/model"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	CreateMessage(text string, senderId, chatId int)error
	GetMessages(chatId int)([]*model.Message, error)
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{
		db: db,
	}
}

func (s *store)CreateMessage(text string, senderId, chatId int)error{
	_, err := s.db.Exec(
		`INSERT INTO messages (sender_id, chat_id, text)
		VALUES ($1, $2, $3)`, senderId, chatId, text,
	)

	return err
}

func (s *store)GetMessages(chatId int)([]*model.Message, error){
	messages := []*model.Message{}

	err := s.db.Select(&messages, `SELECT * FROM messages WHERE chat_id = $1`, chatId)

	return messages, err
}