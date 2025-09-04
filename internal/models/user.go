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

func ChangeName(userName string, ID uuid.UUID) User {
	if userName == "" {
		return User{}
	}
	err := database.DB.Model(&User{}).Where("id = ?", ID).Update("user_name", userName).Error
	if err != nil {
		return User{}
	}
	var u User
	err = database.DB.Model(&User{}).First(&u, "id = ?", ID).Error
	if err != nil {
		return User{}
	}
	return u
}

func DeleteUser(ID uuid.UUID) []User {
	var users []User
	err := database.DB.Delete(&User{}, "id = ?", ID).Error
	if err != nil {
		return nil
	}
	return users
}

func UpdateName(userName string, ID uuid.UUID) []User {
	var users []User
	err := database.DB.Model(&User{}).Where("id = ?", ID).Update("user_name", userName).Error
	if err != nil {
		return nil
	}
	return users
}
