package store

import (
	"check-mate/internal/model"

	"github.com/jmoiron/sqlx"
)

type Store interface {
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
		user_id INTEGER
	   );
	`

	s.db.Exec(query)
}
