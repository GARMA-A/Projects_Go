package middleware

import (
	"errors"
	"net/http"
)

var UnAuthorizedError = errors.New("Invalid token or user name")

func Authrization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		var err error
		if username == "" || token == "" {
			err = errors.New("username or token is missing")
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if token != "12345" {
			http.Error(w, UnAuthorizedError.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
