package formatters

type UserDataFormatter struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type UserSignInFormatter struct {
	Id            int    `json:"id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Authorization string `json:"authorization"`
}
