package dto

type CreateAccountAppDTO struct {
	Email    string `json:"user_email"`
	Account  string `json:"user_account"`
	Password string `json:"user_password"`
	Salt     string `json:"user_salt"`
	Username string `json:"user_name"`
	Phone    string `json:"user_phone"`
}