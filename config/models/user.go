package models

type User struct {
	ID          int64  `json:"id"`
	UserName    string `json:"username"`
	UserSurname string `json:"user_surname"`
	Password    string `json:"password"`
}
