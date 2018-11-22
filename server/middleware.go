package main

import (
	"net/http"
	"strings"
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
