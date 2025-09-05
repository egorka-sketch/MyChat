package main

import (
	"MyChat/config"
	"MyChat/internal/database"
	"MyChat/internal/handlers"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	slog.Info("Запуск чата")
	cfg := config.LoadConfig("config/config.yaml")

	database.InitDB(
		cfg.Database.Host,
		fmt.Sprintf("%d", cfg.Database.Port),
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	log.Println("Сервер запущен")

	http.HandleFunc("/createUser", handlers.CreateUser)
	http.HandleFunc("/postMessage", handlers.PostMessage)
	http.HandleFunc("/getContact", handlers.GetContact)
	http.HandleFunc("/createContact", handlers.CreateContact)
	http.HandleFunc("/getChat", handlers.GetChat)
	http.HandleFunc("/deleteUser", handlers.DeleteUser)
	http.HandleFunc("/deleteContact", handlers.DeleteContact)
	http.HandleFunc("/updateName", handlers.UpdateUser)
	http.HandleFunc("/updateContact", handlers.UpdateContact)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
