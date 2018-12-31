package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// ContentTypeJSONMiddleware is a middleware that appends a JSON header for our APIs
func ContentTypeJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// All APIs should return JSON
		if strings.HasPrefix(path, "/api") {
			w.Header().Add("Content-Type", "application/json")
		}

		next.ServeHTTP(w, r)
	})
}

// JWTAuthentication is a middleware that prevents unauthorized access to the API
func JWTAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			next.ServeHTTP(w, r)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check if correct signing method was used
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return SigningKey, nil
		})

		if !token.Valid || err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}

		if err != nil {
			log.Print(err)
			w.Write([]byte("Token could not be parsed"))
		}

		next.ServeHTTP(w, r)
	})
}
