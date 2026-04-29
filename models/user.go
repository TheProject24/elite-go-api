package models

type RegisterRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	StartingBalance int `json:"starting_balance"`
}
