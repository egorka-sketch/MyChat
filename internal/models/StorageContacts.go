package models

var ContactStorage *StorageContacts

type StorageContacts struct {
	Contacts map[string][]*Contacts
}

func NewStorageContacts() {
	ContactStorage = &StorageContacts{
		Contacts: make(map[string][]*Contacts)}

}

//func (cs *StorageContacts) AddContact(Cont Contacts) {
//	cs.Contacts[Cont.ContactID] = append(cs.Contacts[Cont.ContactID], &Cont)
//}

func (cs *StorageContacts) GetContact(id string) []*Contacts {
	return cs.Contacts[id]
}

func (cs *StorageContacts) DeleteContact(id string) {
	delete(cs.Contacts, id)
}
