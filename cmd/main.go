package main

import (
	"MyChat/internal/database"
	"MyChat/internal/handlers"
	"net/http"
)

func main() {
	database.InitDB()

	http.HandleFunc("/createUser", handlers.CreateUser)
	http.HandleFunc("/postMessage", handlers.PostMessage)
	http.HandleFunc("/getContact", handlers.GetContact)
	http.HandleFunc("/createContact", handlers.CreateContact)
	http.HandleFunc("/GetChat", handlers.GetChat)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
