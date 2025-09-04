package models

type User struct {
	UserId   int    `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	// password string `json:"password"`
}
