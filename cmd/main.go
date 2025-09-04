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
	slog.Info("Запуск чата")

	database.InitDB()

	http.HandleFunc("/createUser", handlers.CreateUser)
	http.HandleFunc("/postMessage", handlers.PostMessage)
	http.HandleFunc("/getContact", handlers.GetContact)
	http.HandleFunc("/createContact", handlers.CreateContact)
	http.HandleFunc("/GetChat", handlers.GetChat)
	http.HandleFunc("/DeleteUser", handlers.DeleteUser)
	http.HandleFunc("/DeleteContact", handlers.DeleteContact)
	http.HandleFunc("/UpdateName", handlers.UpdateUserName)
	http.HandleFunc("/UpdateContact", handlers.UpdateContact)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
