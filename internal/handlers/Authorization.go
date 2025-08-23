package handlers

import "net/http"

type Authorization struct {
	Login         string `json:"Login"`
	Password      string `json:"Password"`
	Success       bool   `json:"Success"`
	StorageAccess string `json:"StorageAccess"`
}

func HandleAuthorization(w http.ResponseWriter, r *http.Request) {
	data := Authorization{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	data.Success = true
}
