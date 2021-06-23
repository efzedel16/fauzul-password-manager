package inputs

type CreatePass struct {
	Web  string `json:"web" binding:"required"`
	Pass string `json:"password" binding:"required"`
}

type UpdatePass struct {
	Web  string `json:"web"`
	Pass string `json:"password"`
}
