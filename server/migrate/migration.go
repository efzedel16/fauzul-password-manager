package migrate

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

type Password struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"user_id"`
	Webs      string    `json:"web"`
	Pass      string    `json:"pass"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
