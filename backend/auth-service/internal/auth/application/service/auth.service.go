package service

import (
	"context"

	httpDto "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/controller/dto"
)

type AuthService interface {
	Login(ctx context.Context, login *httpDto.UserLoginReq) (int, *httpDto.UserLoginRes, error)
	Logout(ctx context.Context, token string) error
	CreateAccount(ctx context.Context, account *httpDto.CreateAccountReq) (int, error)
}