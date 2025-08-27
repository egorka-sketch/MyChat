package main

import (
	"MyChat/internal/database"
	"MyChat/internal/models"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type createContactData struct {
	ContactName string `json:"ContactName"`
	ContactId   string `json:"ContactId"`
	UserID      string `json:"UserID"`
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
		http.Error(w, "NOT CORRECT JSON", http.StatusBadRequest)
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

//	func GetMessage(w http.ResponseWriter, r *http.Request) {
//		if r.Method != "GET" {
//			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//			return
//		}
//
//		recipientID := r.URL.Query().Get("id")
//		if recipientID == "" {
//			http.Error(w, "recipientID is required", http.StatusBadRequest)
//			return
//		}
//
//		msg := models.MessagesStorage.GetMessageByRecipient(recipientID)
//		if len(msg) == 0 {
//			http.Error(w, "No message found", http.StatusNotFound)
//			return
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		err := json.NewEncoder(w).Encode(msg)
//		if err != nil {
//			return
//		}
//	}
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
	}
	receiverID, err := uuid.Parse(data.ReceiverId)
	if err != nil {
		http.Error(w, "RECEIVER ID is not a valid UUID", http.StatusBadRequest)
	}
	mes := models.NewMessage(senderID, receiverID, data.Text)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mes)
	if err != nil {
		http.Error(w, "JSON NOT CORRECTION", http.StatusBadRequest)
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
	contactID, err := uuid.Parse(ContactID)
	ContactName := r.URL.Query().Get("contact_name")
	if ContactName == "" {
		http.Error(w, "contact_name is required", http.StatusBadRequest)
		return
	}
	cont := models.GetContact(ContactName, contactID)

	if cont == nil {
		http.Error(w, "No contact found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cont)
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
}

func main() {
	database.InitDB()
	models.NewMessageStorage()
	models.NewUserStorage()
	models.NewStorageContacts()

	http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/postMessage", PostMessage)
	//http.HandleFunc("/getMessage", GetMessage)
	http.HandleFunc("/getContact", GetContact)
	http.HandleFunc("/createContact", CreateContact)
	http.HandleFunc("/GetChat", GetChat)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
