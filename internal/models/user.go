package models

import (
	"MyChat/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;column:id"`
	UserName string    `gorm:"column:user_name"`
}

func NewUser(userName string) User {
	userId, _ := uuid.NewUUID()
	u := User{
		ID:       userId,
		UserName: userName,
	}
	err := database.DB.Create(&u).Error
	if err != nil {
		return User{}
	}
	return u
}

func (u *User) ChangeName(newUserName string) {
	u.UserName = newUserName
}
