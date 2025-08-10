package models

import "github.com/google/uuid"

var ContactStorage *StorageContacts

type StorageContacts struct {
	Contacts map[uuid.UUID]*Contacts
}

func NewStorageContacts() {
	ContactStorage = &StorageContacts{
		Contacts: make(map[uuid.UUID]*Contacts)}

}
func (cs *StorageContacts) AddContact(Contacts Contacts) {
	cs.Contacts[Contacts.ID] = &Contacts
}

func (cs *StorageContacts) GetContact(id uuid.UUID) *Contacts {
	return cs.Contacts[id]
}

func (cs *StorageContacts) DeleteContact(id uuid.UUID) {
	delete(cs.Contacts, id)
}
