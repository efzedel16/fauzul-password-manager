package inputs

type SignUp struct {
	FullName string `json:"full_name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Pass     string `json:"pass" binding:"required"`
}

type SignIn struct {
	Email string `json:"email" binding:"required,email"`
	Pass  string `json:"pass" binding:"required"`
}

//type UpdateUser struct {
//	FullName string `json:"full_name"`
//	Address  string `json:"address"`
//	Email    string `json:"email"`
//	Pass     string `json:"pass"`
//}
