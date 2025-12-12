package model

type UserModel struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}
