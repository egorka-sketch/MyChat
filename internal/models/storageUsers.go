package models

import "github.com/google/uuid"

var UsersStorage *UserStorage

type UserStorage struct {
	s map[uuid.UUID]*User
}

func NewUserStorage() {
	UsersStorage = &UserStorage{
		s: make(map[uuid.UUID]*User),
	}
}

func (us *UserStorage) AddUser(user User) {
	us.s[user.ID] = &user
}
func (us *UserStorage) GetUser(id uuid.UUID) *User {
	return us.s[id]
}

func (us *UserStorage) RemoveUser(id uuid.UUID) {
	delete(us.s, id)
}
