package entities

import "time"

type User struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	FullName  string     `json:"full_name"`
	Email     string     `json:"email" gorm:"unique"`
	PassHash  string     `json:"pass_hash"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	PassApp   []Password `gorm:"foreignKey:UserId"`
}
