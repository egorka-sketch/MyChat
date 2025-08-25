package models

import (
	"MyChat/internal/database"
	"github.com/google/uuid"
)

type Contacts struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;column:id"`
	ContactName string    `gorm:"column:contact_name"`
	ContactID   uuid.UUID `gorm:"column:contact_id"`
	UserID      uuid.UUID `gorm:"column:user_id"`
}

func NewContact(ContactName string, ContactID uuid.UUID, UserID uuid.UUID) Contacts {
	CId, _ := uuid.NewUUID()
	Cont := Contacts{
		ID:          CId,
		ContactName: ContactName,
		ContactID:   ContactID,
		UserID:      UserID,
	}
	err := database.DB.Create(&Cont).Error
	if err != nil {
		return Contacts{}
	}
	return Cont
}

func GetContact(ContactName string, ContactID uuid.UUID) []Contacts {
	var contacts []Contacts
	err := database.DB.Where("contact_id=? OR contact_name=?", ContactID, ContactName).Find(&contacts).Error
	if err != nil {
		return []Contacts{}
	}
	return contacts
}
