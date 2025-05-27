package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/QuanCters/backend/internal/database"
	"github.com/QuanCters/backend/internal/model"
	"github.com/QuanCters/backend/internal/utils/crypto"
	"github.com/QuanCters/backend/pkg/response"
)

type sUserAdmin struct {
	r *database.Queries
}

func UserAdminImpl(r *database.Queries) *sUserAdmin{
	return &sUserAdmin{
		r: r,
	}
}

func (s *sUserAdmin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	userFound, err := s.r.CheckUserExist(ctx, in.Email)

	if err != nil {
		return response.ErrCodeUserHasExists, err
	}

	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	salt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to generate salt: %w", err)
	}

	hashedPassword := crypto.HashPassword(in.Password, salt)

	account := strings.Split(in.Email, "@")[0]

	params := database.AddUserParams{
		UserAccount:  account,
		UserEmail:    in.Email,
		UserPassword: hashedPassword,
		UserSalt:     salt,
		UserName: 	  in.Name,
		UserPhone:    in.Phone,
	}

	_, err = s.r.AddUser(ctx, params)
	if err != nil {
		return response.ErrCodeInternal, fmt.Errorf("failed to create user: %w", err)
	}

	return response.ErrCodeSuccess, nil
}