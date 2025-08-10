package models

import (
	"github.com/google/uuid"
)

type Contacts struct {
	ID          uuid.UUID
	ContactName string
	ContactID   uuid.UUID
}

func NewContact(ContactName string, ContactID uuid.UUID) Contacts {
	CId, _ := uuid.NewUUID()
	Cont := Contacts{
		ID:          CId,
		ContactName: ContactName,
		ContactID:   ContactID,
	}
	ContactStorage.AddContact(Cont)
	return Cont
}
