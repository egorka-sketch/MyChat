package main

import (
	"MyChat/internal/models"
	"encoding/json"
	"net/http"
)

type createContactData struct {
	ContactName string `json:"ContactName"`
	ContactId   string `json:"ContactId"`
}

type createUserData struct {
	UserName string `json:"username"`
}
type createNewMessage struct {
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Text       string `json:"text"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "POST Only", http.StatusMethodNotAllowed)
		return
	}

	var data createUserData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	user := models.NewUser(data.UserName)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "success", "message": "Created user: " + user.UserName}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	recipientID := r.URL.Query().Get("id")
	if recipientID == "" {
		http.Error(w, "recipientID is required", http.StatusBadRequest)
		return
	}

	msg := models.MessagesStorage.GetMessageByRecipient(recipientID)
	if len(msg) == 0 {
		http.Error(w, "No message found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		return
	}
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
	}

	mes := models.NewMessage(data.SenderId, data.ReceiverId, data.Text)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "success", "message": "Message sending: " + mes.Text}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
func GetContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET", http.StatusMethodNotAllowed)
		return
	}

	ContactID := r.URL.Query().Get("contact_id")
	if ContactID == "" {
		http.Error(w, "contact_id is required", http.StatusBadRequest)
		return
	}

	cont := models.ContactStorage.GetContact(ContactID)
	if cont == nil {
		http.Error(w, "No contact found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(cont)
	if err != nil {
		return
	}
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
	cont := models.NewContact(data.ContactName, data.ContactId)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cont)
	if err != nil {
		return
	}
}

func main() {
	models.NewMessageStorage()
	models.NewUserStorage()
	models.NewStorageContacts()

	http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/postMessage", PostMessage)
	http.HandleFunc("/getMessage", GetMessage)
	http.HandleFunc("/getContact", GetContact)
	http.HandleFunc("/createContact", CreateContact)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
