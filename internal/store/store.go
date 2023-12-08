package store

import "github.com/jmoiron/sqlx"

type Store interface {
	CreateMessage(text string, senderId, receiverId int)error
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{
		db: db,
	}
}

func (s *store)CreateMessage(text string, senderId, receiverId int)error{
	_, err := s.db.DB.Exec(
		`INSERT INTO messages (sender_id, receiver_id, text)
		VALUES ($1, $2, $3)`, senderId, receiverId, text,
	)

	return err
}
