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

func NewContact(ContactName string, ContactID, UserID uuid.UUID) (Contacts, error) {
	cont := Contacts{
		ID:          ContactID,
		ContactName: ContactName,
		ContactID:   ContactID,
		UserID:      UserID,
	}
	err := database.DB.Create(&cont).Error
	if err != nil {
		return Contacts{}, err
	}
	return cont, nil
}

func GetContact(UserID string) []Contacts {
	var contacts []Contacts
	err := database.DB.Where("user_id = ?", UserID).Find(&contacts).Error
	if err != nil {
		return []Contacts{}
	}
	return contacts
}

func DeleteContact(ContactId uuid.UUID) []Contacts {
	var contacts []Contacts
	err := database.DB.Where("contact_id", ContactId).Delete(&contacts).Error
	if err != nil {
		return []Contacts{}
	}
	return contacts
}
