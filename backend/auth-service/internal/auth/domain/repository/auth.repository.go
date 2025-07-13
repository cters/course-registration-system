package repository

import (
	"context"

	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/domain/entity"
)

type AuthRepository interface {
	CheckUserExist(ctx context.Context, email string) (bool, error)
	CreateAccountWithTx(ctx context.Context, account *entity.Account) (int32, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.Account, error)
}