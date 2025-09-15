package handlers

import (
	"MyChat/internal/models"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type createUserData struct {
	UserName string `json:"username"`
	ID       string `json:"id"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	response := map[string]string{"status": "success", "userId": user.ID.String()}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "PUT Only", http.StatusMethodNotAllowed)
		return
	}
	var data createUserData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "NOT CORRECT JSON", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(data.ID)
	UpUser := models.ChangeName(data.UserName, id)
	if err != nil {
		http.Error(w, "Failed to update", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(UpUser)
	if err != nil {
		return
	}
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Delete Only", http.StatusMethodNotAllowed)
		return
	}

	Id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	DeleteId := models.DeleteUser(Id)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(DeleteId)
	if err != nil {
		return
	}
}

func GetIdUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET Only", http.StatusMethodNotAllowed)
		return
	}
	userName := r.URL.Query().Get("username")
	if userName == "" {
		http.Error(w, "username parameter is required", http.StatusBadRequest)
		return
	}
	GetId := models.GetUsers(userName)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(GetId)
	if err != nil {
		return
	}
}
