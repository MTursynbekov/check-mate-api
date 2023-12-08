package model

type Message struct {
	ID int `json:"id" db:"id"`
	SenderID int `json:"senderId" db:"sender_id"`
	ReceiverID int `json:"receiverId" db:"receiver_id"`
	Text string  `json:"text" db:"text"`
}