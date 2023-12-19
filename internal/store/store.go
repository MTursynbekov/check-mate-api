package store

import (
	"check-mate/internal/model"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	CreateMessage(text string, senderId, chatId int) error
	GetMessages(chatId int) ([]*model.Message, error)
	CreateChat(firstMemberId, secondMemberId int) error
	GetChat(id int) ([]*model.Chat, error)
	CreateContact(contact *model.Contact) (int, error)
	GetContacts(userId int) ([]*model.Contact, error)
	GetContact(userId, contactId int) (*model.Contact, error)
	UpdateContact(contact *model.Contact) error
	DeleteContact(contactId int) error
	Migrate()
	CreateUser(user *model.User) (uint, error)
	GetUser(username string) (*model.User, error)
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{
		db: db,
	}
}

func (s *store) Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR UNIQUE,
		password VARCHAR,
		phone VARCHAR,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP
	);
	
	CREATE TABLE IF NOT EXISTS chats (
		id SERIAL PRIMARY KEY,
		first_member_id INTEGER,
		second_member_id INTEGER
	   );
	   
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		chat_id INTEGER REFERENCES chats(id) ON DELETE CASCADE,
		sender_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		content TEXT NOT NULL
	   );
	   
	CREATE TABLE IF NOT EXISTS contacts (
		id SERIAL PRIMARY KEY,
		surname TEXT NOT NULL,
		name TEXT NOT NULL,
		relationship TEXT NOT NULL,
		user_id INTEGER,
		reminder_time TIMESTAMP
		birthday TIMESTAMP
	   );
	`

	s.db.Exec(query)
}

func (s *store) CreateMessage(text string, senderId, chatId int) error {
	_, err := s.db.Exec(
		`INSERT INTO messages (sender_id, chat_id, content)
		VALUES ($1, $2, $3)`, senderId, chatId, text,
	)

	return err
}

func (s *store) GetMessages(chatId int) ([]*model.Message, error) {
	messages := []*model.Message{}

	err := s.db.Select(&messages, `SELECT * FROM messages WHERE chat_id = $1`, chatId)

	return messages, err
}

func (s *store) CreateChat(firstMemberId, secondMemberId int) error {
	_, err := s.db.Exec(
		`INSERT INTO chats (first_member_id, second_member_id)
		VALUES ($1, $2)`, firstMemberId, secondMemberId,
	)

	return err
}

func (s *store) GetChat(id int) ([]*model.Chat, error) {
	chats := []*model.Chat{}

	err := s.db.Select(&chats, `SELECT * FROM chats WHERE first_member_id = $1`, id)

	return chats, err
}
