package models

import (
	"github.com/google/uuid"
)

type Contacts struct {
	ID          uuid.UUID
	ContactName string
	ContactID   string
}

func NewContact(ContactName string, ContactID string) Contacts {
	CId, _ := uuid.NewUUID()
	Cont := Contacts{
		ID:          CId,
		ContactName: ContactName,
		ContactID:   ContactID,
	}
	ContactStorage.AddContact(Cont)
	return Cont
}
