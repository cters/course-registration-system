package dto

// CREATE ACCOUNT DTOs
type CreateAccountReq struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
	Phone string `json:"phone" validate:"omitempty,min=10,max=15"`
	Role string `json:"role" validate:"required,oneof=admin user"`
	Name string `json:"name" validate:"required,min=3,max=50"`
}

type CreateAccountRes struct {
	UserID int   `json:"user_id"`
}

// LOGIN DTOs
type UserLoginReq struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserLoginRes struct {
	UserID   int32  `json:"id"`
	UserAccount string `json:"account"`
	Username string `json:"name"`
	UserEmail    string `json:"email"`
	UserPhone string `json:"phone"`
	Token string `json:"token"`
} 