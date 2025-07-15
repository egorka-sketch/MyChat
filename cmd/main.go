package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "host=localhost user=MyChat password=7782 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Ошибка подключения к базе данных")

	}
	db.AutoMigrate(&User{}, &Contact{}, &Message{})
}
func main() {
	fmt.Println("Hello World")
}
