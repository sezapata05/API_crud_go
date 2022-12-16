package api

import (
	"net/http"

	"github.com/google/uuid"
)

var users = map[string]string{"user1": "password1", "user2": "password2"}

func requestIDHanler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := r.Header.Get("x-Request-ID")

		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}

		w.Header().Set("x-Request-ID", requestID)

		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := r.Header.Get("Authorization")

		if users[user] == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("Authorization", user)

		next.ServeHTTP(w, r)
	})
}
