package models

type User struct {
	full_name string `json:"full_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
