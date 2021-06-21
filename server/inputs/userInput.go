package inputs

type InputUserSignUp struct {
  FullName string `json:"full_name"`
  Address string `json:"address"`
  Email string `json:"email"`
  Password string `json:"password"`
}