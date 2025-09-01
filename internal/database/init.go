package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

var DB *gorm.DB

func InitDB() {
	slog.Info("Подключение базы данных")

	dsn := "host=localhost user=mychat password=7782 dbname=mychatdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("не удалось подключиться к базе данных: %v", err)
	}
	slog.Info("Подключение успешно")
}
