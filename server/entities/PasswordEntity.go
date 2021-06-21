package entities

type PassApp struct {
  Id        string `json:"id"`
  UserId    []User `json:"user_id"`
  No        string `json:"no"`
  Website   string `json:"website"`
  CreatedAt string `json:"created_at"`
  UpdateAt  string `json:"update_at"`
}