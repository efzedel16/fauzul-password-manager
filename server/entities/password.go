package entities

import "time"

type Password struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"user_id"`
	Web       string    `json:"web"`
	Pass      string    `json:"pass"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
