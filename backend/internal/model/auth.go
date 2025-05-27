package model

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

type RegisterOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type LogoutInput struct {
	Token string `json:"token"`
	ID    string `json:"user_id"`
}

type LogoutOutput struct {
	Message string `json:"message"`
}