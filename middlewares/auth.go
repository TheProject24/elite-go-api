package middlewares

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "my-secret-token" {
			log.Println("Security Alert: Blocked unauthorized request")
			http.Error(w, "Access Denied: Invalid Token", http.StatusUnauthorized)
			return
		}

		log.Println(("Token verified. Passing to handler. . . "))
		next(w, r)
	}
}