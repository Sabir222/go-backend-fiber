package models

import "time"

// User represents the structure of our resource
type User struct {
	ID           string    `json:"id"`
	FullName     string    `json:"full_name"`
	PasswordHash string    `json:"password_hash"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Role         string    `json:"role"`
}
