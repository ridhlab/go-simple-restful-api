package user

type CreateUserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	UserId   int     `json:"user_id"`
	Fullname *string `json:"fullname"`
	Email    *string `json:"email"`
}
