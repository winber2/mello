package main

import (
	"crypto/rand"
	"log"
	"mello/server/mongo"
	"mello/server/mongo/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// SigningKey is a randomly generated key to create/validate the JWT
var SigningKey = getRandomKey()

func getRandomKey() []byte {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

// GetAuthToken checks if valid credentials were given and returns a token
func GetAuthToken(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	user, err := models.GetUserFromCredentials(username, password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mongo.ToJSON(err))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"admin":    true,
	})

	// Sign the token with our secret
	tokenString, _ := token.SignedString(SigningKey)

	w.Write([]byte(tokenString))
}
