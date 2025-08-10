package models

import "github.com/google/uuid"

type User struct {
	//gorm.Model
	ID       uuid.UUID
	UserName string
}

func NewUser(userName string) User {
	userId, _ := uuid.NewUUID()
	u := User{
		ID:       userId,
		UserName: userName,
	}
	UsersStorage.AddUser(u)
	return u
}

func (u *User) ChangeName(newUserName string) {
	u.UserName = newUserName
}
