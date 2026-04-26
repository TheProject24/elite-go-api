package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func healthCheckerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := Response{
		Status: "success",
		Message: "Taiwo's Elite Go API is live",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/health", healthCheckerHandler)

	log.Println("Server booting up on port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server crashed:", err)
	}
}