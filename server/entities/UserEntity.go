package entities

type User struct {
  Id        int `json:"id"`
  FullName  string `json:"full_name"`
  Email     string `json:"email"`
  PasswordHash string `json:"password"`
  Address   string `json:"address"`
  PassApp   []PassApp `json:"pass_app"`
  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
}