package Middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var JwtKey []byte

func SetJWTKey(key []byte) {
	JwtKey = key
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		Authorize := r.Header.Get("Authorization")
		if Authorize == "" {
			http.Error(w, "No Authorization header", http.StatusUnauthorized)
			return
		}
		parts := strings.Split(Authorize, " ")

		if len(parts) != 2 {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return JwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
