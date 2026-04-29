package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TheProject24/elite-api/models"
)

	func RegisterUser(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req models.RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		bgCtx := context.WithoutCancel(r.Context())

		go func() {
			err := sendWelcomeEmail(bgCtx, req.Email)
			if err != nil {
				log.Printf("Failed to send email to %s: %v", req.Email, err)
			}
		}()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User Registered!"})
	}

