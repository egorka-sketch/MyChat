package models

import (
	"github.com/google/uuid"
)

type Message struct {
	//gorm.Model
	ID          uuid.UUID
	SenderID    string
	RecipientID string
	Text        string
}

func NewMessage(senderID string, receiverID string, text string) Message {
	messageId, _ := uuid.NewUUID()
	m := Message{
		ID:          messageId,
		SenderID:    senderID,
		RecipientID: receiverID,
		Text:        text,
	}
	MessagesStorage.AddMessage(m)
	return m
}
