package entities

import "time"

type User struct {
	Id           int 			 `json:"id"`
	FullName     string 	 `json:"full_name"`
	Email        string 	 `json:"email"`
	PasswordHash string 	 `json:"password_hash"`
	Address      string 	 `json:"address"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	PassApp      []Password
}