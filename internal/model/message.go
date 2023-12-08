package model

type Message struct {
	ID int `json:"id" db:"id"`
	SenderID int `json:"senderId" db:"sender_id"`
	ChatID int `json:"chatId" db:"chat_id"`
	Text string  `json:"text" db:"text"`
}