package entities

type User struct {
	Id           int
	FullName     string
	Email        string
	PasswordHash string
	Address      string
	CreatedAt    string
	UpdatedAt    string
	//PassApp      []PassApp
}