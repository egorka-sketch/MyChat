package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

var DB *gorm.DB

func InitDB(host, port, user, pass, name string) {
	slog.Info("Подключение базы данных")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, pass, name, port,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("не удалось подключиться к базе данных: %v", err)
	}
	slog.Info("Подключение успешно")
}
