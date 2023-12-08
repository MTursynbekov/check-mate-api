package model

type Message struct {
	ID int `json:"id" db:"id"`
	SenderID int `json:"senderId" db:"sender_id"`
	ChatID int `json:"chatId" db:"chat_id"`
	Content string  `json:"content" db:"content"`
}