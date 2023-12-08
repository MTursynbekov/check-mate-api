package model

type Contact struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Surname      string `json:"surname" db:"surname"`
	Relationship string `json:"relationship" db:"relationship"`
	UserID       int    `json:"userId" db:"user_id"`
	Reminder string `json:"reminder" db:"reminder"`
}
