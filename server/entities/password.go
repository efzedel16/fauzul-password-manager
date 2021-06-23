package entities

import "time"

type Password struct {
  Id        int `json:"id"`
  UserId    int `json:"user_id"`
  Website   string `json:"web"`
  Password  string `json:"password"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}