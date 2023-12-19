package store

import (
	"check-mate/internal/model"
	"database/sql"
)

func (s *store) CreateContact(contact *model.Contact) error {
	query := `
	INSERT INTO contacts (name, surname, relationship, user_id, reminder_time)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(query, contact.Name, contact.Surname, contact.Relationship, contact.UserID, contact.ReminderTime)

	return err
}

func (s *store) GetContacts(userId int) ([]*model.Contact, error) {
	contacts := make([]*model.Contact, 0)
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

func (s *store) UpdateContact(contact *model.Contact) error {
	query := `
	UPDATE contacts
	SET name=$1,
		surname=$2,
		relationship=$3,
		remider_time=$4
	WHERE id=$5`

	_, err := s.db.Exec(query, contact.Name, contact.Surname, contact.Relationship, contact.ReminderTime)

	return err
}

func (s *store) DeleteContact(contactId int) error {
	query := `
	DELETE FROM contacts
	WHERE id=$1`

	_, err := s.db.Exec(query, contactId)
	return err
}
