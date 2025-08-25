package models

import (
	"MyChat/internal/database"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;column:id"`
	SenderID    uuid.UUID `gorm:"column:sender_id"`
	RecipientID uuid.UUID `gorm:"column:recipient_id"`
	Text        string    `gorm:"column:text"`
	Timestamp   time.Time `gorm:"Times"`
}

func NewMessage(senderID uuid.UUID, receiverID uuid.UUID, text string) Message {
	messageId, _ := uuid.NewUUID()
	m := Message{
		ID:          messageId,
		SenderID:    senderID,
		RecipientID: receiverID,
		Text:        text,
	}
	err := database.DB.Create(&m).Error
	if err != nil {
		return Message{}
	}
	return m
}
func GetAllMessage(UserId string) []Message {
	var messages []Message
	err := database.DB.Where("sender_id = ? OR recipient_id = ?", UserId, UserId).Order("times ASC").Find(&messages).Error
	if err != nil {
		return []Message{}
	}
	return messages
}
