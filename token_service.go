package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secure-devops-key")

// Token generate function
func createToken(username string) (string, error) {

	claims := jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

// Login API
func loginAPI(w http.ResponseWriter, r *http.Request) {

	user := r.URL.Query().Get("user")

	if user == "" {
		http.Error(w, "User required", http.StatusBadRequest)
		return
	}

	token, err := createToken(user)

	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Token: %s", token)
}

// Secure endpoint
func protectedAPI(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		http.Error(w, "Token missing", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Secure API Access Granted")
}

func main() {

	http.HandleFunc("/login", loginAPI)
	http.HandleFunc("/secure", protectedAPI)

	fmt.Println("Token service running on port 7072")

	http.ListenAndServe(":7072", nil)
}