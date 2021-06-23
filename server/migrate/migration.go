package migrate

import "time"

type User struct {
	Id           int        `json:"id" gorm:"primaryKey"`
	FullName     string     `json:"full_name"`
	Email        string     `json:"email" gorm:"unique"`
	PasswordHash string     `json:"password_hash"`
	Address      string     `json:"address"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    time.Time  `json:"deleted_at" gorm:"index"`
	PassApp      []Password `gorm:"foreignKey:userId"`
}

type Password struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"user_id"`
	Website   string    `json:"web"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}
