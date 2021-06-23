package entities

import "time"

type Password struct {
  Id        int       `json:"id" gorm:"primaryKey"`
  UserId    int       `json:"user_id"`
  Website   string    `json:"web"`
  Password  string    `json:"password"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}