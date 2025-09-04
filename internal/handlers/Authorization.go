package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

var jwtKey []byte

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

func GenerateJWT(UserID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = UserID
	claims["admin"] = true
	return token.SignedString(jwtKey)
}
