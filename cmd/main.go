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
