package main

import (
	"MyChat/internal/database"
	"MyChat/internal/handlers"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	slog.Info("Запуск приложения MyChat")

	database.InitDB()

	http.HandleFunc("/createUser", handlers.CreateUser)
	http.HandleFunc("/postMessage", handlers.PostMessage)
	http.HandleFunc("/getContact", handlers.GetContact)
	http.HandleFunc("/createContact", handlers.CreateContact)
	http.HandleFunc("/GetChat", handlers.GetChat)
	http.HandleFunc("/UpdateUserName", handlers.UpdateUser)
	http.HandleFunc("/DeleteUser", handlers.DeleteUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
