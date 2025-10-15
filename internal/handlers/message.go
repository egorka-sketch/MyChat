package handlers

import (
	"MyChat/internal/models"
	"encoding/json"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type createNewMessage struct {
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Text       string `json:"text"`
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	msg := models.GetAllMessage(userID)

	if len(msg) == 0 {
		http.Error(w, "No messages found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		return
	}
	slog.Info("GetChat", "user", userID)
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST", http.StatusMethodNotAllowed)
		return
	}

	var data createNewMessage

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "JSON NOT CORRECTION", http.StatusBadRequest)
		return
	}

	senderID, err := uuid.Parse(data.SenderId)
	if err != nil {
		http.Error(w, "SENDER ID is not a valid UUID", http.StatusBadRequest)
		return
	}

	receiverID, err := uuid.Parse(data.ReceiverId)
	if err != nil {
		http.Error(w, "RECEIVER ID is not a valid UUID", http.StatusBadRequest)
		return
	}

	mes := models.NewMessage(senderID, receiverID, data.Text)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(mes)
	if err != nil {
		http.Error(w, "JSON NOT CORRECTION", http.StatusBadRequest)
		return
	}
	slog.Info("PostMessage", "user", senderID)
}
