package handlers

import (
	"MyChat/internal/models"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type createContactData struct {
	ContactName string `json:"ContactName"`
	ContactId   string `json:"ContactId"`
	UserID      string `json:"UserID"`
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET", http.StatusMethodNotAllowed)
		return
	}
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	cont := models.GetContact(userID)

	if cont == nil {
		http.Error(w, "No contact found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(cont)
	if err != nil {
		return
	}
	slog.Info(fmt.Sprint(cont))
	return
}
func CreateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST", http.StatusMethodNotAllowed)
		return
	}
	var data createContactData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "JSON NOT CORRECTION", http.StatusBadRequest)
	}
	contactID, err := uuid.Parse(data.ContactId)
	if err != nil {
		return
	}
	UserId, err := uuid.Parse(data.UserID)
	if err != nil {
		return
	}
	cont := models.NewContact(data.ContactName, contactID, UserId)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cont)
	if err != nil {
		return
	}

	slog.Info(fmt.Sprint("Create Contact"))
}
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Only DELETE", http.StatusMethodNotAllowed)
		return
	}

	contactID, err := uuid.Parse(r.URL.Query().Get("contact_id"))
	if err != nil {
		return
	}

	deleteContact := models.DeleteContact(contactID)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(deleteContact)
	if err != nil {
		return
	}
	slog.Info(fmt.Sprint("Contact Deleted"))
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Only PUT", http.StatusMethodNotAllowed)
		return
	}
	var data createContactData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "JSON NOT CORRECTION", http.StatusBadRequest)
	}
	contactID, err := uuid.Parse(r.URL.Query().Get("contact_id"))
	if err != nil {
		return
	}
	userID, err := uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		return
	}

	UpContact := models.UpdateContact(data.ContactName, contactID, userID)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(UpContact)
	if err != nil {
		return
	}
}
