package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=mychat password=7782 dbname=mychatdb port=5432 sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("не удалось подключиться к базе данных: %v", err)
	}

	DB = db
	log.Println("Подключение к базе данных успешно!")
}
