package main

import (
	"MyChat/config"
	"MyChat/internal/Middleware"
	"MyChat/internal/database"
	"MyChat/internal/handlers"
	"fmt"
	"log"
	"log/slog"
	. "net/http"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	slog.Info("Запуск чата")
	cfg := config.LoadConfig("config/config.yaml")

	handlers.SetJWTKey([]byte(cfg.JWT.Key))
	Middleware.SetJWTKey([]byte(cfg.JWT.Key))

	database.InitDB(
		cfg.Database.Host,
		fmt.Sprintf("%d", cfg.Database.Port),
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	log.Println("Сервер запущен")

	HandleFunc("/register", handlers.Register)
	HandleFunc("/login", handlers.Login)

	HandleFunc("/createUser", Middleware.JWT(handlers.CreateUser))
	HandleFunc("/postMessage", Middleware.JWT(handlers.PostMessage))
	HandleFunc("/getContact", Middleware.JWT(handlers.GetContact))
	HandleFunc("/createContact", Middleware.JWT(handlers.CreateContact))
	HandleFunc("/getChat", Middleware.JWT(handlers.GetChat))
	HandleFunc("/deleteUser", Middleware.JWT(handlers.DeleteUser))
	HandleFunc("/deleteContact", Middleware.JWT(handlers.DeleteContact))
	HandleFunc("/updateName", Middleware.JWT(handlers.UpdateUser))
	HandleFunc("/getIdUser", Middleware.JWT(handlers.GetIdUser))

	err := ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
