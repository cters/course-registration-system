package model

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
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

type LogoutOutput struct {
	Message string `json:"message"`
}

type UserOutput struct {
	UserID      int32  `json:"user_id"`
	UserAccount string `json:"user_account"`
	UserEmail   string `json:"user_email"`
	UserName    string `json:"user_name"`
	UserPhone   string `json:"user_phone"`
}