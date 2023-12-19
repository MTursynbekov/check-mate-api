package store

import (
	"check-mate/internal/model"
	"database/sql"
)

func (s *store) CreateContact(contact *model.Contact) (int, error) {
	query := `
	INSERT INTO contacts (name, surname, relationship, user_id, reminder_time)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(query, contact.Name, contact.Surname, contact.Relationship, contact.UserID, contact.ReminderTime)
	if err != nil {
		return -1, err
	}
	
	query = `
	SELECT id FROM contacts
	ORDER BY id DESC
	LIMIT 1
	`

	var id int

	row := s.db.QueryRow(query)
	row.Scan(&id)

	return id, err
}

func (s *store) GetContacts(userId int) ([]*model.Contact, error) {
	var contacts []*model.Contact
	query := `
	SELECT * FROM contacts
	WHERE user_id = $1`

	err := s.db.Select(&contacts, query, userId)

	return contacts, err
}

func (s *store) GetContact(userId, contactId int) (*model.Contact, error) {
	var contact model.Contact
	query := `
	SELECT * FROM contacts
	WHERE user_id = $1 and id = $2`

	err := s.db.Get(&contact, query, userId, contactId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &contact, err
}
