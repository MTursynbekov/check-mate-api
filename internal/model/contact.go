package model

import "time"

type Contact struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Surname      string    `json:"surname" db:"surname"`
	Relationship string    `json:"relationship" db:"relationship"`
	UserID       int       `json:"userId" db:"user_id"`
	ReminderTime time.Time `json:"reminder" db:"reminder_time"`
}
