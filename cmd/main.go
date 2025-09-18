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

	HandleFunc("/register", handlers.RegisterHandler)
	HandleFunc("/login", handlers.LoginHandler)

	HandleFunc("/createUser", Middleware.JWTMiddleware(handlers.CreateUser))
	HandleFunc("/postMessage", Middleware.JWTMiddleware(handlers.PostMessage))
	HandleFunc("/getContact", Middleware.JWTMiddleware(handlers.GetContact))
	HandleFunc("/createContact", Middleware.JWTMiddleware(handlers.CreateContact))
	HandleFunc("/getChat", Middleware.JWTMiddleware(handlers.GetChat))
	HandleFunc("/deleteUser", Middleware.JWTMiddleware(handlers.DeleteUser))
	HandleFunc("/deleteContact", Middleware.JWTMiddleware(handlers.DeleteContact))
	HandleFunc("/updateName", Middleware.JWTMiddleware(handlers.UpdateUser))
	HandleFunc("/getIdUser", Middleware.JWTMiddleware(handlers.GetIdUser))

	err := ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
