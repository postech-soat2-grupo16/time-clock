package user

type User struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Registration string `json:"registration"`
	Password     string `json:"password"`
}
