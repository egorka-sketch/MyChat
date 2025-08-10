package main

import (
	"MyChat/internal/models"
	"encoding/json"
	"net/http"
)

type createUserData struct {
	UserName string `json:"username"`
}
type createNewMessage struct {
	SenderId   string `json:"sender_id"`   //переделать тип юид на стринг. в жсоне не бывает типа юид
	ReceiverId string `json:"receiver_id"` //переделать тип юид на стринг. в жсоне не бывает типа юид
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

func main() {
	models.NewMessageStorage()
	models.NewUserStorage()

	http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/postMessage", PostMessage)
	http.HandleFunc("/getMessage", GetMessage)
	http.HandleFunc("/sendMessage", PostMessage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
